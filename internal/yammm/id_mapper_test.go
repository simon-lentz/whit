package yammm_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/yammm"
)

func Test_IDMapper(t *testing.T) {
	tt := testutils.NewTester(t)
	model := `schema "testing" type Subject {  }`
	// instance has one id being only local, one id being both local and global, and one being only global.
	instance := `{ "Subjects": [
		{"id": "$$local"},
		{"id": "$$local2"},
	 	{"id": "$$:kDEyuWY2vJ3JMFv42Us4eV:local2"},
	 	{"id": "$$:kDEyuWY2vJ3JMFv42Us4eV:local3"}
		]}
		`
	inst := makeJzonInstance(t.Name(), instance)

	ctx, ic := makeContext(t.Name(), model, true)
	tt.CheckFalse(ic.HasFatal() || ic.HasErrors())

	idMapper := yammm.NewIDMapper()
	replacementMap, err := idMapper.Map(ctx, inst)
	tt.CheckNotError(err)
	// should have entries for "$$local" and "$$local2"
	uid1, ok := replacementMap["$$local"]
	var nullUUID uuid.UUID
	tt.CheckNotEqual(uid1, nullUUID)
	tt.CheckTrue(ok)
	uid2, ok := replacementMap["$$local2"]
	tt.CheckTrue(ok)
	uid3, ok := replacementMap["$$:kDEyuWY2vJ3JMFv42Us4eV:local2"]
	tt.CheckTrue(ok)
	// Assert that the local uuid for $$local2 is the same as for the global $$:...:local2.
	tt.CheckEqual(uid2, uid3)
	// Assert that the "only global" resulted in entry for its local
	_, ok = replacementMap["$$local3"]
	tt.CheckTrue(ok)
}

// Association properties and Pk properties (can skip from type values as they are resolved).
