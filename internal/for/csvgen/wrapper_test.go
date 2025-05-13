package csvgen_test

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/for/csvgen"
	"github.com/wyrth-io/whit/internal/xray"
)

func TestWrapperRow(t *testing.T) {
	tt := testutils.NewTester(t)
	model := `schema "test"
	type Person {
		name String primary
		age Integer required
		popularity Float required
		isImaginary Boolean required
		-->OWNS (many) Car { since Date }
	}
	type Car {
		regNbr String primary
	}
	`
	ctx, _ := makeContext(t.Name(), model, true)

	genModel := `{
	"type": "Person",
	"property_map": {"name": "NAME", "age": "AGE", "popularity": "RATING", "isImaginary": "IMAGINARY"},
	"association_map": {
		"OWNS_Cars": { 
			"property_map": { "since": "OWNERSHIP_SINCE"},
			"where": { "regNbr": "REGISTRATION_NUMBER"}
		}
	}
}`
	csvData := `NAME,AGE,REGISTRATION_NUMBER,OWNERSHIP_SINCE,RATING,IMAGINARY
Fred Flintstone,42,ABC123,2023-01-02,3.14,true
Henry Ford,63,T1,1908-01-02,0.2,false
`
	file := makeTmpFile(csvData)
	defer os.Remove(file.Name()) //nolint:errcheck

	var gm csvgen.Genmodel
	err := json.Unmarshal([]byte(genModel), &gm)
	tt.CheckNotError(err)

	f, err := os.Open(file.Name())
	tt.CheckNotError(err)
	reader := csv.NewReader(f)

	graph, err := csvgen.NewGraph(ctx, &gm, reader)
	tt.CheckNotError(err)
	tt.CheckNotNil(graph)
	w := xray.NewWrapper(graph)
	tt.CheckTrue(w.IsObject())
	tt.CheckEqual([]string{"People"}, w.FeatureNames())

	instancesFeature := w.Feature("People")
	tt.CheckTruef(instancesFeature.IsSlice(), "csv top wrapper should report it is a Slice")
	tt.CheckEqual(2, instancesFeature.Len())

	row1 := instancesFeature.FeatureAtIndex(0)
	tt.CheckTruef(row1.IsObject(), "row wrapper should report it is an Object")
	tt.CheckEqual("Fred Flintstone", row1.Value("name"))
	tt.CheckEqual(42, row1.Value("age"))
	tt.CheckEqual(3.14, row1.Value("popularity"))
	tt.CheckEqual(true, row1.Value("isImaginary"))

	row2 := instancesFeature.FeatureAtIndex(1)
	tt.CheckTruef(row1.IsObject(), "row wrapper should report it is an Object")
	tt.CheckEqual("Henry Ford", row2.Value("name"))
	tt.CheckEqual(63, row2.Value("age"))
	tt.CheckEqual(0.2, row2.Value("popularity"))
	tt.CheckEqual(false, row2.Value("isImaginary"))

	// Associations on row 1
	ownedCars := row1.Feature("OWNS_Cars")
	tt.CheckTruef(ownedCars.IsSlice(), "should be slice of objects since relation is many")
	tt.CheckEqual(1, ownedCars.Len())
	oneOwned := ownedCars.FeatureAtIndex(0)
	tt.CheckTruef(oneOwned.IsObject(), "the relationship should be an object")
	since := oneOwned.Value("since")
	tt.CheckEqual("2023-01-02", since)
	where := oneOwned.Feature("Where")
	tt.CheckTruef(where.IsObject(), "should be object with primary keys")
	tt.CheckEqual(1, where.Len()) // should be one property
	tt.CheckEqual("ABC123", where.Value("regNbr"))

	// Association on row 2
	ownedCars = row2.Feature("OWNS_Cars")
	tt.CheckTruef(ownedCars.IsSlice(), "should be slice of objects since relation is many")
	tt.CheckEqual(1, ownedCars.Len())
	oneOwned = ownedCars.FeatureAtIndex(0)
	tt.CheckTruef(oneOwned.IsObject(), "the relationship should be an object")
	since = oneOwned.Value("since")
	tt.CheckEqual("1908-01-02", since)
	where = oneOwned.Feature("Where")
	tt.CheckTruef(where.IsObject(), "should be object with primary keys")
	tt.CheckEqual(1, where.Len()) // should be one property
	tt.CheckEqual("T1", where.Value("regNbr"))
}
