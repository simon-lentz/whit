package csvgen

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/wyrth-io/whit/internal/pio"
	"github.com/wyrth-io/whit/internal/utils"
)

// MissingValuesError is an error indicating that there were missing values in one or more
// columns in a CSV.
type MissingValuesError struct {
	// Cols contains the names of the columns that have no values.
	Cols []string
}

func (e *MissingValuesError) Error() string {
	return "the following columns are empty: " + strings.Join(
		utils.Map(e.Cols, func(s string) string { return fmt.Sprintf("%q", s) }),
		", ",
	)
}

// HeaderToYammm produces Yammm output to the given writer. The csv header columns are given in
// firstRow, and their corresponding mapped "camel" names are given in camels. The reader is
// the csv reader used to obtain rows. The main purpose of this is to find the base data type
// per column by searching for the first value in every column. If at the end there are columns
// still remaining, the special NoValuesError are returned with a slice of the csv headers for
// the columns having no values. The search stops as soon as all columns have got their type.
// The 'sourceName' parameter is used in error messages and in the yammm output. The ´schemaName' and
// 'typeName' parameters can be left empty in which case the name of the schema and type is derived
// from the 'sourceName' (which is typically given as a filepath).
func HeaderToYammm(sourceName string, writer *pio.Writer, reader *csv.Reader, schemaName, typeName string) error {
	firstRow, camels, err := GetFirstRowAndCamels(reader)
	if err != nil {
		return err
	}
	cols := make([]int, len(firstRow))
	dataTypes := make([]string, len(firstRow))

	// Create set of columns to get the type for in column order
	for i := range firstRow {
		cols[i] = i
	}
	lineNbr := 0
	for {
		// if there are no more columns needing a type.
		if len(cols) < 1 {
			break
		}
		// get the next row
		lineNbr++
		sampleRow, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("could not read row %d of '%s': %s", lineNbr, sourceName, err.Error())
		}
		reexamine := []int{}
		for _, col := range cols {
			val := sampleRow[col]
			switch {
			case IsBoolean(val) || strings.HasSuffix(firstRow[col], "?"):
				dataTypes[col] = "Boolean"
			case IsFloat(val):
				dataTypes[col] = "Float"
			case IsInteger(val):
				dataTypes[col] = "Integer"
			default:
				if val == "" {
					reexamine = append(reexamine, col)
				}
				dataTypes[col] = "String"
			}
		}
		cols = reexamine
	}
	// Produce the output
	schemaName, typeName = GetDerivedNames(sourceName, schemaName, typeName)
	writer.FormatLn(`/* schema derived from: %s */`, sourceName)
	writer.FormatLn(`schema "%s"`, schemaName)
	writer2 := writer.FormatLn("type %s {", typeName).Indented()
	for i := range firstRow {
		writer2.FormatLn("/* %s */", firstRow[i])
		writer2.FormatLn("%s %s", camels[i], dataTypes[i])
	}
	writer.FormatLn("}")

	if len(cols) > 0 {
		names := make([]string, len(cols))
		for i := range cols {
			names[i] = firstRow[cols[i]]
		}
		return &MissingValuesError{Cols: names}
	}
	return nil
}

// HeaderToMapping produces a csv genmodel with mappings from property names to column names read
// from the CSV header in the given reader. If the given type name is "" the name of the type is
// derived from the sourcePath.
func HeaderToMapping(sourcePath string, writer *pio.Writer, reader *csv.Reader, typeName string) error {
	_, typeName = GetDerivedNames(sourcePath, "", typeName)

	gm := Genmodel{}
	gm.AssociationMap = map[string]*AssocModel{}
	gm.Typename = typeName
	firstRow, camels, err := GetFirstRowAndCamels(reader)
	if err != nil {
		return err
	}
	propMap := make(map[string]string, len(firstRow))
	for i := range firstRow {
		propMap[camels[i]] = firstRow[i]
	}
	gm.PropertyMap = propMap
	bytes, err := json.MarshalIndent(gm, "", "  ")
	if err != nil {
		return fmt.Errorf("could not generate json for genmodel: %s", err.Error())
	}
	writer.Println(string(bytes))
	return nil
}

// GetFirstRowAndCamels reads the header row from the csv reader and returns it and
// a slice where all column names have been converted to lower camelCase form. An
// error is returned if these could not be produced.
func GetFirstRowAndCamels(reader *csv.Reader) (firstRow, camels []string, err error) {
	// Read and process the header
	firstRow, err = reader.Read()
	if err != nil {
		return nil, nil, err
	}
	camels = make([]string, len(firstRow))
	camelSet := utils.NewSet[string]()
	for i := range firstRow {
		c := utils.ToLowerCamel(firstRow[i])
		if c == "" {
			return nil, nil, fmt.Errorf("column '%d' has empty header", i)
		}
		if !camelSet.Add(c) {
			return nil, nil, fmt.Errorf("mapped column '%s' is not unique", c)
		}
		camels[i] = c
	}
	return firstRow, camels, nil
}

// GetDerivedNames returns a schemaName and a typeName derived from the given sourcePath if
// the givenSchemaName or givenTypeName are empty.
func GetDerivedNames(sourcePath, givenSchemaName, givenTypeName string) (schemaName, typeName string) {
	basePath := filepath.Base(sourcePath)
	derivedName := strings.TrimSuffix(basePath, filepath.Ext(basePath))
	derivedTypeName := utils.ToUpperCamel(derivedName)
	if givenSchemaName != "" {
		derivedName = givenSchemaName
	}
	if givenTypeName != "" {
		derivedTypeName = givenTypeName
	}
	return derivedName, derivedTypeName
}
