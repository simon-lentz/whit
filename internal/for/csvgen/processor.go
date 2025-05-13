package csvgen

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/wyrth-io/whit/internal/pio"
	"github.com/wyrth-io/whit/internal/tc"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/yammm"
)

// Process processes a csv file given by name, using a compleated and validated yammm model,
// a genmodel for property mapping and writes the generated JSON output to the given writer.
func Process(csvFile string, ctx yammm.Context, gm Genmodel, writer *pio.Writer) error { //nolint:gocognit
	// Open and read the file line by line
	f, err := os.Open(csvFile) //nolint:gosec
	if err != nil {
		return err
	}
	reader := csv.NewReader(f)

	// Read and process the header
	firstLine, err := reader.Read()
	if err != nil {
		return err
	}
	colMap := make(map[string]int, len(firstLine))
	for i := range firstLine {
		colMap[firstLine[i]] = i
	}
	colNames := utils.NewSet(utils.Keys(colMap)...)

	// Get the type from the Genmodel, it must exist in the context.
	targetType := ctx.LookupType(gm.Typename)
	if targetType == nil {
		return fmt.Errorf("genmodel typename %s is not the name of a type in the given model", gm.Typename)
	}
	// Get all the properties from the type and make a map.
	allProps := targetType.AllProperties()
	propMap := make(map[string]*yammm.Property, len(allProps))
	for _, p := range allProps {
		propMap[p.Name] = p
	}
	// Get all the associations from the type and make a map
	allAssocs := targetType.AllAssociations()
	assocMap := make(map[string]*yammm.Association, len(allAssocs))
	for _, a := range allAssocs {
		assocMap[a.PropertyName(ctx)] = a
	}

	// PROPERTIES.
	// Assert that genmodel only mentions target type properties.
	mapped := utils.NewSet(utils.Keys(gm.PropertyMap)...)
	existing := utils.NewSet(utils.Keys(propMap)...)
	extra := mapped.Diff(existing)
	if extra.Size() != 0 {
		return fmt.Errorf("genmodel type '%s' does not have the mapped properties: %v", targetType.Name, extra.Slices())
	}
	// Assert that the genmodel column names all exist in the csv header.
	mapped = utils.NewSet[string](utils.Values(gm.PropertyMap)...)
	extra = mapped.Diff(colNames)
	if extra.Size() != 0 {
		return fmt.Errorf("csv file '%s' does not have the mapped columns: %v", csvFile, extra.Slices())
	}

	// ASSOCIATIONS.
	// Assert that genmodel only mentions target type associations.
	mapped = utils.NewSet(utils.Keys(gm.AssociationMap)...)
	existing = utils.NewSet(utils.Keys(assocMap)...)
	extra = mapped.Diff(existing)
	if extra.Size() != 0 {
		return fmt.Errorf("genmodel type '%s' does not have the mapped associations: %v", targetType.Name, extra.Slices())
	}
	for aName, aModel := range gm.AssociationMap {
		a := assocMap[aName]
		// All mapped properties must be properties of the association.
		mapped = utils.NewSet(utils.Keys(aModel.Properties)...)
		existing = utils.NewSetFrom(a.Properties, func(p *yammm.Property) string { return p.Name })
		extra = mapped.Diff(existing)
		if extra.Size() != 0 {
			return fmt.Errorf("association '%s' does not have mapped properties: %v", aName, extra.Slices())
		}
		// All mapped columns must exist in the csv
		mapped = utils.NewSet(utils.Values(aModel.Properties)...)
		extra = mapped.Diff(colNames)
		if extra.Size() != 0 {
			return fmt.Errorf("csv file '%s' does not have the mapped columns: %v", csvFile, extra.Slices())
		}

		// All primary keys in the Where must exist as primary keys of target type and as columns.
		mapped = utils.NewSet(utils.Keys(aModel.Where)...)
		aTarget := ctx.LookupType(a.To) // Model is valid, so cannot fail.
		// Where must have either ID or all other primary keys
		if mapped.Contains("id") {
			if mapped.Size() != 1 {
				return fmt.Errorf("where map using 'id' should be the only key used, got: %v", mapped.Slices())
			}
		} else {
			pksNoID := utils.NewSet(
				utils.Map(aTarget.AllPrimaryKeys(), func(p *yammm.Property) string { return p.Name })...,
			).Remove("id")
			if !pksNoID.Equal(mapped) {
				return fmt.Errorf("all non id primary keys were not given, required: %v, got: %v",
					pksNoID.Slices(), mapped.Slices())
			}
		}
		// Make sure mapped primary key columns exist.
		mapped = utils.NewSet(utils.Values(aModel.Where)...)
		extra = mapped.Diff(colNames)
		if extra.Size() != 0 {
			return fmt.Errorf("mapped primary key columns %v does not match columns in csv file", extra.Slices())
		}
	}

	// Phew, validatation of gendata done.
	// Start document:
	writer.FormatLn("{")
	out := writer.Indented()
	out.FormatLn(`"%s":`, targetType.PluralName) // the [] around instances is done as one marshal at the end

	lineNbr := 0 // LineNbr is used for error messages. Zero is for the header.
	instances := []map[string]any{}
	for {
		lineNbr++
		record, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("error reading csv line %d: %s", lineNbr, err.Error())
		}
		instance := make(map[string]any, len(gm.PropertyMap)+len(gm.AssociationMap))
		for pn, cn := range gm.PropertyMap {
			v := record[colMap[cn]] // unconverted string
			v2, err := mapDataType(lineNbr, v, pn, propMap)
			if err != nil {
				return err
			}
			if v2 != nil {
				instance[pn] = v2
			}
		}
		for pn, am := range gm.AssociationMap {
			// get a map of the modeled association's properties
			a := assocMap[pn]
			// TODO: Wasteful, this builds this map for every line in the csv. Could be cached.
			aPropsMap := make(map[string]*yammm.Property, len(a.Properties))
			for _, p := range a.Properties {
				aPropsMap[p.Name] = p
			}

			// create the output structure
			where := map[string]any{}
			aProps := map[string]any{"Where": where}
			// map association properties
			for apn, cn := range am.Properties {
				v := record[colMap[cn]]
				v2, err := mapDataType(lineNbr, v, apn, aPropsMap)
				if err != nil {
					return err
				}
				if v2 != nil {
					aProps[apn] = v2
				}
			}
			// map association primary keys
			// TODO: This is wasteful, should be cached.
			targetType := ctx.LookupType(a.To)
			allPks := targetType.AllPrimaryKeys()
			pkMap := make(map[string]*yammm.Property, len(allPks))
			for _, p := range allPks {
				pkMap[p.Name] = p
			}
			for pkn, cn := range am.Where {
				v := record[colMap[cn]]
				v2, err := mapDataType(lineNbr, v, pkn, pkMap)
				if err != nil {
					return err
				}
				if v2 != nil {
					where[pkn] = v2
				}
			}
			if a.Many {
				instance[pn] = []any{aProps}
			} else {
				instance[pn] = aProps
			}
		}
		instances = append(instances, instance)
	}
	// Serialize all instances to json (wasteful, but handles the "no comma at the end" issue).
	bytes, err := json.MarshalIndent(instances, out.Spaces(), "  ")
	if err != nil {
		return fmt.Errorf("could not marshal json (internal error): %s", err.Error())
	}
	writer.Println(string(bytes))
	writer.FormatLn(`}`)

	return nil
}
func mapDataType(lineNbr int, v string, pn string, propMap map[string]*yammm.Property) (any, error) {
	if v == "" {
		// missing value (Lint complains but I don't want a sentinel error here).
		return nil, nil //nolint:nilnil
	}
	switch propMap[pn].BaseType().Kind() {
	case tc.StringKind:
		return v, nil
	case tc.IntKind:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("csv data lind %d for property '%s', value is not a valid integer, got: %s",
				lineNbr, pn, v)
		}
		return i, nil
	case tc.FloatKind:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, fmt.Errorf("csv data line %d for property '%s', value is not a valid float, got: %s",
				lineNbr, pn, v)
		}
		return f, nil
	case tc.BoolKind:
		// Be lenient
		var b bool
		switch strings.ToLower(v) {
		case "yes", "y", "true", "t", "1":
			b = true
		case "no", "n", "false", "f", "0", "-1", "":
			b = false
		default:
			return nil, fmt.Errorf("csv data line %d for property '%s', value is not a boolean, got: %s",
				lineNbr, pn, v)
		}
		return b, nil
	default:
		return nil, fmt.Errorf("internal error (should not happen) - unknown kind")
	}
}
