package tc_test

import (
	"math"
	"testing"

	"github.com/google/uuid"
	"github.com/hlindberg/testutils"
	"github.com/lithammer/shortuuid/v4"
	"github.com/pkg/errors"
	"github.com/wyrth-io/whit/internal/tc"
)

func Test_IntegerTypeChecker(t *testing.T) {
	assertOk(t, -1, []string{"Integer", "", ""})
	assertOk(t, 0, []string{"Integer", "", ""})
	assertOk(t, 1, []string{"Integer", "", ""})
	assertOk(t, math.MinInt64, []string{"Integer", "", ""})
	assertOk(t, math.MaxInt64, []string{"Integer", "", ""})
	assertMessage(t, -1, "Integer value -1 is < min 0", []string{"Integer", "0", ""})
	assertMessage(t, 2, "Integer value 2 is > max 1", []string{"Integer", "0", "1"})

	assertMessage(t, nil, "nil value is not an Integer", []string{"Integer", "", ""})
	assertMessage(t, 0.0, "value is not an Integer", []string{"Integer", "", ""})
	assertMessage(t, false, "value is not an Integer", []string{"Integer", "", ""})
	assertMessage(t, "rosebud", "value is not an Integer", []string{"Integer", "", ""})
}

func Test_FloatTypeChecker(t *testing.T) {
	assertOk(t, -1.0, []string{"Float", "", ""})
	assertOk(t, 0.0, []string{"Float", "", ""})
	assertOk(t, 1.0, []string{"Float", "", ""})
	assertOk(t, -math.MaxFloat64, []string{"Float", "", ""})
	assertOk(t, math.MaxFloat64, []string{"Float", "", ""})
	assertMessage(t, -1.0, "Float value -1 is < min 0", []string{"Float", "0.0", ""})
	assertMessage(t, 2.0, "Float value 2 is > max 1", []string{"Float", "0.0", "1.0"})

	// accepts integers as Float
	assertOk(t, 0, []string{"Float", "", ""})
	assertOk(t, 1, []string{"Float", "", ""})
	assertOk(t, -1, []string{"Float", "", ""})

	assertMessage(t, nil, "nil value is not a Float", []string{"Float", "", ""})
	assertMessage(t, false, "value is not a Float", []string{"Float", "", ""})
	assertMessage(t, "rosebud", "value is not a Float", []string{"Float", "", ""})
}

func Test_StringTypeChecker(t *testing.T) {
	assertOk(t, "", []string{"String", "", ""})
	assertOk(t, "a", []string{"String", "", ""})
	assertOk(t, "abc", []string{"String", "", ""})
	assertMessage(t, "", "String length 0 is shorter than min allowed 1", []string{"String", "1", ""})
	assertMessage(t, "ab", "String length 2 is longer than max allowed 1", []string{"String", "0", "1"})

	assertMessage(t, nil, "nil value is not a String", []string{"String", "", ""})
	assertMessage(t, 0.0, "value is not a String", []string{"String", "", ""})
	assertMessage(t, false, "value is not a String", []string{"String", "", ""})
	assertMessage(t, 1, "value is not a String", []string{"String", "", ""})
}

func Test_BooleanTypeChecker(t *testing.T) {
	assertOk(t, false, []string{"Boolean"})
	assertOk(t, true, []string{"Boolean"})

	assertMessage(t, nil, "nil value is not a Boolean", []string{"Boolean"})
	assertMessage(t, 0.0, "value is not a Boolean", []string{"Boolean"})
	assertMessage(t, "rosebud", "value is not a Boolean", []string{"Boolean"})
	assertMessage(t, 1, "value is not a Boolean", []string{"Boolean"})
}
func Test_DateTypeChecker(t *testing.T) {
	assertOk(t, "2023-06-29", []string{"Date"})

	assertMessage(t, "notadate", "value 'notadate' does not match Date format '2006-01-02'", []string{"Date"})
	assertMessage(t, "2023-13-01", "value '2023-13-01' does not match Date format '2006-01-02'", []string{"Date"})
	assertMessage(t, "2023-01-40", "value '2023-01-40' does not match Date format '2006-01-02'", []string{"Date"})

	assertMessage(t, nil, "nil value is not a Date", []string{"Date"})
	assertMessage(t, 0.0, "value is not a Date", []string{"Date"})
	assertMessage(t, false, "value is not a Date", []string{"Date"})
	assertMessage(t, 1, "value is not a Date", []string{"Date"})
}

func Test_TimestampTypeChecker(t *testing.T) {
	assertOk(t, "2006-01-02T15:04:05Z", []string{"Timestamp"})

	assertMessage(t, "notadate",
		`value does not match Timestamp format : parsing time "notadate" as "2006-01-02T15:04:05Z07:00": cannot parse "notadate" as "2006"`,
		[]string{"Timestamp"})
	assertMessage(t, "2023-13-01",
		`value does not match Timestamp format : parsing time "2023-13-01": month out of range`,
		[]string{"Timestamp"})
	assertMessage(t, "2023-01-40T00:00:00Z",
		`value does not match Timestamp format : parsing time "2023-01-40T00:00:00Z": day out of range`,
		[]string{"Timestamp"})

	assertMessage(t, nil, "nil value is not a Timestamp", []string{"Timestamp"})
	assertMessage(t, 0.0, "value is not a Timestamp", []string{"Timestamp"})
	assertMessage(t, false, "value is not a Timestamp", []string{"Timestamp"})
	assertMessage(t, 1, "value is not a Timestamp", []string{"Timestamp"})
}

func Test_EnumTypeChecker(t *testing.T) {
	rgb := []string{"Enum", "red", "green", "blue"} // note order
	assertOk(t, "blue", rgb)
	assertOk(t, "green", rgb)
	assertOk(t, "red", rgb)

	// note sorted order in output
	assertMessage(t, "black", `String value 'black' does not match Enum["blue", "green", "red"]`, rgb)
	assertMessage(t, "", `String value '' does not match Enum["blue", "green", "red"]`, rgb)

	assertMessage(t, nil, "nil value is not an Enum", rgb)
	assertMessage(t, 0.0, "value is not a String, cannot match an Enum", rgb)
	assertMessage(t, false, "value is not a String, cannot match an Enum", rgb)
	assertMessage(t, 1, "value is not a String, cannot match an Enum", rgb)
}
func Test_UUIDTypeChecker(t *testing.T) {
	tt := testutils.NewTester(t)
	uid := []string{"UUID"} // note order
	assertOk(t, "$$local", uid)
	base57 := "kDEyuWY2vJ3JMFv42Us4eV"
	assertOk(t, base57, uid)
	assertOk(t, "$$:"+base57+":2", uid)
	uuid, err := shortuuid.DefaultEncoder.Decode(base57)
	tt.CheckNil(err)
	uuidV4String := uuid.String()
	assertOk(t, uuidV4String, uid)

	assertMessage(t, nil, "nil value is not an UUID", uid)
	assertMessage(t, 0.0, "value is not a string based UUID", uid)
	assertMessage(t, false, "value is not a string based UUID", uid)
	assertMessage(t, 1, "value is not a string based UUID", uid)
}
func Test_PatternTypeChecker(t *testing.T) {
	rgb := []string{"Pattern", "red", "green", "blue"} // note order
	assertOk(t, "blue", rgb)
	assertOk(t, "green", rgb)
	assertOk(t, "red", rgb)

	// note sorted order in output
	assertMessage(t, "black", `String value 'black' does not match Pattern["blue", "green", "red"]`, rgb)
	assertMessage(t, "", `String value '' does not match Pattern["blue", "green", "red"]`, rgb)

	assertMessage(t, nil, "nil value is not a String, cannot match a Pattern", rgb)
	assertMessage(t, 0.0, "value is not a String, cannot match a Pattern", rgb)
	assertMessage(t, false, "value is not a String, cannot match a Pattern", rgb)
	assertMessage(t, 1, "value is not a String, cannot match a Pattern", rgb)
}

func Test_SpacevecTypeChecker(t *testing.T) {
	vector3 := []string{"Spacevector", "3"}
	assertOk(t, []float32{1, 2, 3}, vector3)
	assertOk(t, []float64{1, 2, 3}, vector3)

	assertMessage(t, nil, "nil value is not a Spacevector", vector3)
	assertMessage(t, []float32{1, 2}, "length of given spacevector must be 3: got 2", vector3)
	assertMessage(t, []float32{1, 2, 3, 4}, "length of given spacevector must be 3: got 4", vector3)
	assertMessage(t, []float64{1, 2, math.MaxFloat64}, "value is outside range of allowe spacevector float32 values", vector3)
	assertMessage(t, []float64{1, 2, -math.MaxFloat64}, "value is outside range of allowe spacevector float32 values", vector3)
}
func Test_BaseType_of_TypeChecker(t *testing.T) {
	assertBaseType(t, tc.StringKind, []string{"String", "", ""})
	assertBaseType(t, tc.IntKind, []string{"Integer", "", ""})
	assertBaseType(t, tc.FloatKind, []string{"Float", "", ""})
	assertBaseType(t, tc.BoolKind, []string{"Boolean"})
	assertBaseType(t, tc.StringKind, []string{"Date"})
	assertBaseType(t, tc.StringKind, []string{"UUID"})
	assertBaseType(t, tc.StringKind, []string{"Timestamp"})
	assertBaseType(t, tc.StringKind, []string{"Enum", "a", "b"})
	assertBaseType(t, tc.StringKind, []string{"Pattern", "a", "b"})
	assertBaseType(t, tc.SpacevectorKind, []string{"Spacevector", "1536"})
}
func Test_TypeString_of_TypeChecker(t *testing.T) {
	assertTypeString(t, []string{"String", "", ""})
	assertTypeString(t, []string{"Integer", "", ""})
	assertTypeString(t, []string{"Float", "", ""})
	assertTypeString(t, []string{"Boolean"})
	assertTypeString(t, []string{"Date"})
	assertTypeString(t, []string{"UUID"})
	assertTypeString(t, []string{"Timestamp"})
	assertTypeString(t, []string{"Enum", "a", "b"})
	assertTypeString(t, []string{"Pattern", "a", "b"})
	assertTypeString(t, []string{"Spacevector", "1536"})
}

func Test_UUIDCheckerGet(t *testing.T) {
	tt := testutils.NewTester(t)
	uuidTC := tc.NewTypeChecker([]string{"UUID"})
	uid, err := uuidTC.GetUUID("$$local")
	var nullUUID uuid.UUID
	tt.CheckEqual(nullUUID, uid)
	tt.CheckTrue(errors.Is(err, tc.ErrRequiresResolution))

	base57 := "kDEyuWY2vJ3JMFv42Us4eV"
	uid, err = uuidTC.GetUUID(base57)
	tt.CheckNotError(err)
	uuid, err := shortuuid.DefaultEncoder.Decode(base57)
	tt.CheckNotError(err)
	tt.CheckEqual(uuid, uid)

	uuidV4String := uuid.String()
	uid, err = uuidTC.GetUUID(uuidV4String)
	tt.CheckNotError(err)
	tt.CheckEqual(uuid, uid)

	globalShort := "$$:" + base57 + ":2"
	uid, err = uuidTC.GetUUID(globalShort)
	tt.CheckNotError(err)
	tt.CheckEqual(uuid, uid)
}

func assertBaseType(t *testing.T, bt tc.Kind, instructions []string) {
	t.Helper()
	tt := testutils.NewTester(t)
	tc := tc.NewTypeChecker(instructions)
	tt.CheckNotNil(tc)
	tt.CheckTruef(bt == tc.BaseType().Kind(), "")
}
func assertOk(t *testing.T, v any, instructions []string) {
	t.Helper()
	tt := testutils.NewTester(t)
	tc := tc.NewTypeChecker(instructions)
	tt.CheckNotNil(tc)
	ok, msg := tc.Check(v)
	tt.CheckTrue(ok)
	tt.CheckEqual("", msg)
}
func assertMessage(t *testing.T, v any, expected string, instructions []string) {
	t.Helper()
	tt := testutils.NewTester(t)
	tc := tc.NewTypeChecker(instructions)
	tt.CheckNotNil(tc)
	ok, msg := tc.Check(v)
	tt.CheckFalse(ok)
	tt.CheckEqual(expected, msg)
}
func assertTypeString(t *testing.T, instructions []string) {
	t.Helper()
	tt := testutils.NewTester(t)
	tc := tc.NewTypeChecker(instructions)
	tt.CheckNotNil(tc)
	label := tc.TypeString()
	tt.CheckEqual(instructions[0], label)
}
