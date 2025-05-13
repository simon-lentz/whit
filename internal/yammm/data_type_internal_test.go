package yammm

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hlindberg/testutils"
	"github.com/lithammer/shortuuid/v4"
	"github.com/wyrth-io/whit/internal/validation"
)

func Test_DataType_validate(t *testing.T) {
	tt := testutils.NewTester(t)
	ctx := NewContext()
	ic := validation.NewIssueCollector()
	dt := &DataType{Name: "x", Constraint: []string{"Integer"}}

	dt.validate(ctx, ic)
	tt.CheckEqual(0, ic.Count())

	// must start with lower case name
	ic = validation.NewIssueCollector()
	dt = &DataType{Name: "X", Constraint: []string{"Integer"}}
	dt.validate(ctx, ic)
	tt.CheckEqual(1, ic.Count())
	ms := validation.MessageSet(ic)
	tt.CheckStringSlicesEqual([]string{"invalid data type name 'X' - data type name must start with lower case 'a'-'z'"}, ms.Slices())

	// name cannot be empty
	ic = validation.NewIssueCollector()
	dt = &DataType{Name: "", Constraint: []string{"Integer"}}
	dt.validate(ctx, ic)
	tt.CheckEqual(1, ic.Count())
	ms = validation.MessageSet(ic)
	tt.CheckStringSlicesEqual([]string{"data type has empty name"}, ms.Slices())

	// cannot have error in constraint
	ic = validation.NewIssueCollector()
	dt = &DataType{Name: "x", Constraint: []string{}}
	dt.validate(ctx, ic)
	tt.CheckEqual(1, ic.Count())
	ms = validation.MessageSet(ic)
	tt.CheckStringSlicesEqual([]string{"invalid constraint for data type 'x': []"}, ms.Slices())
}

type MyUUID uuid.UUID
type TestStruct struct {
	ID MyUUID `json:"id,omitempty"`
}

func (x *MyUUID) MarshalText() (text []byte, err error) {
	str := uuid.UUID(*x).String()
	return []byte(str), nil
}
func (x *MyUUID) UnmarshalText(text []byte) error {
	uid := uuid.MustParse(string(text))
	*x = MyUUID(uid)
	return nil
}
func Test_MarshalUUID(t *testing.T) {
	tt := testutils.NewTester(t)
	base57 := "kDEyuWY2vJ3JMFv42Us4eV"
	uid, err := shortuuid.DefaultEncoder.Decode(base57)
	tt.CheckNil(err)
	data := TestStruct{ID: MyUUID(uid)}
	bytes, err := json.Marshal(data)
	tt.CheckNil(err)
	fmt.Printf("output: %s", bytes)
	// tt.CheckTrue(false)
	var target MyUUID
	err = json.Unmarshal([]byte(`"ed2ab159-cc3d-4236-97fd-4f2c938db14c"`), &target)
	tt.CheckNotError(err)
	tt.CheckEqual("ed2ab159-cc3d-4236-97fd-4f2c938db14c", uuid.UUID(target).String())
}
