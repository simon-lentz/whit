package yammm_test

import (
	"fmt"
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/xray"
	"github.com/wyrth-io/whit/internal/yammm"
)

type testListener struct {
	yammm.BaseGraphListener
	count  int
	events []string
}

func (tl *testListener) EnterGraph(_ yammm.Context, _ xray.Wrapper) {
	tl.count++
	tl.events = append(tl.events, fmt.Sprintf("%d EnterGraph", tl.count))
}
func (tl *testListener) EnterType(_ yammm.Context, _ *yammm.Type) {
	tl.count++
	tl.events = append(tl.events, fmt.Sprintf("%d EnterType", tl.count))
}
func (tl *testListener) EnterInstance(_ yammm.Context, _ *yammm.Type, _ xray.Wrapper) {
	tl.count++
	tl.events = append(tl.events, fmt.Sprintf("%d EnterInstance", tl.count))
}
func (tl *testListener) ExitInstance(_ yammm.Context, _ *yammm.Type, _ xray.Wrapper) {
	tl.count++
	tl.events = append(tl.events, fmt.Sprintf("%d ExitInstance", tl.count))
}

func (tl *testListener) OnProperties(_ yammm.Context, _ *yammm.Type, _ map[string]any) {
	tl.count++
	tl.events = append(tl.events, fmt.Sprintf("%d OnProperties", tl.count))
}

func (tl *testListener) ExitType(_ yammm.Context, _ *yammm.Type) {
	tl.count++
	tl.events = append(tl.events, fmt.Sprintf("%d ExitType", tl.count))
}

func (tl *testListener) ExitGraph(_ yammm.Context, _ xray.Wrapper) {
	tl.count++
	tl.events = append(tl.events, fmt.Sprintf("%d ExitGraph", tl.count))
}
func (tl *testListener) EnterAssociation(_ yammm.Context, _ *yammm.Type, _ *yammm.Association, _ xray.Wrapper) {
	tl.count++
	tl.events = append(tl.events, fmt.Sprintf("%d EnterAssociation", tl.count))
}
func (tl *testListener) ExitAssociation(_ yammm.Context, _ *yammm.Type, _ *yammm.Association, _ xray.Wrapper) {
	tl.count++
	tl.events = append(tl.events, fmt.Sprintf("%d ExitAssociation", tl.count))
}
func (tl *testListener) EnterComposition(_ yammm.Context, _ *yammm.Type, _ *yammm.Composition, _ xray.Wrapper) {
	tl.count++
	tl.events = append(tl.events, fmt.Sprintf("%d EnterAssociation", tl.count))
}
func (tl *testListener) ExitComposition(_ yammm.Context, _ *yammm.Type, _ *yammm.Composition, _ xray.Wrapper) {
	tl.count++
	tl.events = append(tl.events, fmt.Sprintf("%d ExitAssociation", tl.count))
}
func (tl *testListener) OnEdge(_ yammm.Context, _ *yammm.Association, _ map[string]any, _ *yammm.Type,
	_ map[string]any, _ *yammm.Type, _ map[string]any) {
	tl.count++
	tl.events = append(tl.events, fmt.Sprintf("%d OnEdge", tl.count))
}
func (tl *testListener) OnCompositionEdge(_ yammm.Context, _ *yammm.Composition, _ *yammm.Type,
	_ map[string]any, _ *yammm.Type, _ map[string]any) {
	tl.count++
	tl.events = append(tl.events, fmt.Sprintf("%d OnCompositionEdge", tl.count))
}

func Test_Listener(t *testing.T) {
	tt := testutils.NewTester(t)

	model := `schema "testing" type Subject { it String }`
	instance := `{ "Subjects": [ {"it": "red"}, {"it": "blue"} ] }`
	inst := makeJzonInstance(t.Name(), instance)

	ctx, ic := makeContext(t.Name(), model, true)
	tt.CheckFalse(ic.HasFatal() || ic.HasErrors())

	var tl testListener

	walker := yammm.NewGraphWalker(ctx, &tl)
	walker.Walk(inst)

	tt.CheckNotNil(utils.Find(tl.events, func(s string) bool { return s == "1 EnterGraph" }))
	tt.CheckNotNil(utils.Find(tl.events, func(s string) bool { return s == "2 EnterType" }))
	tt.CheckNotNil(utils.Find(tl.events, func(s string) bool { return s == "3 EnterInstance" }))
	tt.CheckNotNil(utils.Find(tl.events, func(s string) bool { return s == "4 OnProperties" }))
	tt.CheckNotNil(utils.Find(tl.events, func(s string) bool { return s == "5 ExitInstance" }))
	tt.CheckNotNil(utils.Find(tl.events, func(s string) bool { return s == "6 EnterInstance" }))
	tt.CheckNotNil(utils.Find(tl.events, func(s string) bool { return s == "7 OnProperties" }))
	tt.CheckNotNil(utils.Find(tl.events, func(s string) bool { return s == "8 ExitInstance" }))
	tt.CheckNotNil(utils.Find(tl.events, func(s string) bool { return s == "9 ExitType" }))
	tt.CheckNotNil(utils.Find(tl.events, func(s string) bool { return s == "10 ExitGraph" }))
	tt.CheckEqual(10, tl.count)
}

func Test_ListenerEnterExitGraph(t *testing.T) {
	tt := testutils.NewTester(t)

	model := `schema "testing" type Subject { it String }`
	instance := `{ "Subjects": [ {"it": "blue"} ] }`
	inst := makeJzonInstance(t.Name(), instance)

	ctx, ic := makeContext(t.Name(), model, true)
	tt.CheckFalse(ic.HasFatal() || ic.HasErrors())

	tl := &yammm.PluggableListener{
		FEnterGraph: func(ctx yammm.Context, data xray.Wrapper) {
			tt.CheckNotNil(utils.Find(data.FeatureNames(), func(s string) bool { return s == "Subjects" }))
		},
		FExitGraph: func(ctx yammm.Context, data xray.Wrapper) {
			tt.CheckNotNil(utils.Find(data.FeatureNames(), func(s string) bool { return s == "Subjects" }))
		},
	}

	walker := yammm.NewGraphWalker(ctx, tl)
	walker.Walk(inst)
}
func Test_ListenerEnterExitType(t *testing.T) {
	tt := testutils.NewTester(t)

	model := `schema "testing" type Subject { it String }`
	instance := `{ "Subjects": [ {"it": "blue"} ] }`
	inst := makeJzonInstance(t.Name(), instance)

	ctx, ic := makeContext(t.Name(), model, true)
	tt.CheckFalse(ic.HasFatal() || ic.HasErrors())

	tl := &yammm.PluggableListener{
		FEnterType: func(ctx yammm.Context, t *yammm.Type) {
			tt.CheckEqual("Subject", t.Name)
		},
		FExitType: func(ctx yammm.Context, t *yammm.Type) {
			tt.CheckEqual("Subject", t.Name)
		},
	}

	walker := yammm.NewGraphWalker(ctx, tl)
	walker.Walk(inst)
}
func Test_ListenerEnterExitInstance(t *testing.T) {
	tt := testutils.NewTester(t)

	model := `schema "testing" type Subject { it String }`
	instance := `{ "Subjects": [ {"it": "blue"} ] }`
	inst := makeJzonInstance(t.Name(), instance)

	ctx, ic := makeContext(t.Name(), model, true)
	tt.CheckFalse(ic.HasFatal() || ic.HasErrors())

	tl := &yammm.PluggableListener{
		FEnterInstance: func(ctx yammm.Context, t *yammm.Type, data xray.Wrapper) {
			tt.CheckEqual("Subject", t.Name)
			tt.CheckEqual("blue", data.Value("it"))
		},
		FExitInstance: func(ctx yammm.Context, t *yammm.Type, data xray.Wrapper) {
			tt.CheckEqual("Subject", t.Name)
			tt.CheckEqual("blue", data.Value("it"))
		},
	}

	walker := yammm.NewGraphWalker(ctx, tl)
	walker.Walk(inst)
}
func Test_ListenerOnProperties(t *testing.T) {
	tt := testutils.NewTester(t)

	model := `schema "testing" type Subject { it String x Integer }`
	instance := `{ "Subjects": [ {"id": "$$1", "it": "blue", "x": 42 } ] }`
	inst := makeJzonInstance(t.Name(), instance)

	ctx, ic := makeContext(t.Name(), model, true)
	tt.CheckFalse(ic.HasFatal() || ic.HasErrors())

	tl := &yammm.PluggableListener{
		FOnProperties: func(ctx yammm.Context, t *yammm.Type, propMap map[string]any) {
			tt.CheckEqual(3, len(propMap))
			tt.CheckEqual("$$1", propMap["id"])
			tt.CheckEqual("blue", propMap["it"])
			tt.CheckEqual(42, propMap["x"])
		},
	}
	walker := yammm.NewGraphWalker(ctx, tl)
	walker.Walk(inst)
}
func Test_ListenerEnterExitAssociation(t *testing.T) {
	tt := testutils.NewTester(t)

	model := `schema "testing"
	type Car {
		regNbr String primary
		--> MADE_BY (one) CarMaker
	}
	type CarMaker {
		name String primary
	}
	`
	instance := `{
		"Cars": [
			{ 	"regnbr": "ABC123",
		  		"MADE_BY": { "Where": { "name": "Fiat" } }
			}
		],
		"CarMakers": [
			{ "name": "Fiat" }
		]
	}`
	inst := makeJzonInstance(t.Name(), instance)

	ctx, ic := makeContext(t.Name(), model, true)
	tt.CheckFalse(ic.HasFatal() || ic.HasErrors())

	tl := &yammm.PluggableListener{
		FEnterAssociation: func(ctx yammm.Context, t *yammm.Type, a *yammm.Association, data xray.Wrapper) {
			tt.CheckEqual("MADE_BY", a.Name)
			tt.CheckEqual("Fiat", data.Feature("Where").Value("name"))
		},
		FExitAssociation: func(ctx yammm.Context, t *yammm.Type, a *yammm.Association, data xray.Wrapper) {
			tt.CheckEqual("MADE_BY", a.Name)
			tt.CheckEqual("Fiat", data.Feature("Where").Value("name"))
		},
	}
	walker := yammm.NewGraphWalker(ctx, tl)
	walker.Walk(inst)
}
func Test_ListenerOnEdgeOne(t *testing.T) {
	tt := testutils.NewTester(t)

	model := `schema "testing"
	type Car {
		regNbr String primary
		--> MADE_BY (one) CarMaker
	}
	type CarMaker {
		name String primary
	}
	`
	instance := `{
		"Cars": [
			{ 	"regNbr": "ABC123",
				"id": "$$1",
		  		"MADE_BY_CarMaker": { "Where": { "name": "Fiat" } }
			}
		],
		"CarMakers": [
			{ "name": "Fiat" }
		]
	}`
	inst := makeJzonInstance(t.Name(), instance)

	ctx, ic := makeContext(t.Name(), model, true)
	tt.CheckFalse(ic.HasFatal() || ic.HasErrors())

	tl := &yammm.PluggableListener{
		FOnEdge: func(ctx yammm.Context, a *yammm.Association, assocPropmap map[string]any,
			fromType *yammm.Type, fromPks map[string]any, toType *yammm.Type, toPks map[string]any,
		) {
			tt.CheckEqual("Car", fromType.Name)
			tt.CheckEqual("CarMaker", toType.Name)
			// from and to PKs only have 'id' !! BUG
			tt.CheckEqual("ABC123", fromPks["regNbr"])
			tt.CheckEqual("Fiat", toPks["name"])
		},
	}
	walker := yammm.NewGraphWalker(ctx, tl)
	walker.Walk(inst)
}

func Test_ListenerOnEdgeMany(t *testing.T) {
	tt := testutils.NewTester(t)

	model := `schema "testing"
	type Car {
		regNbr String primary
		--> MADE_BY (many) CarMaker { x Integer}
	}
	type CarMaker {
		name String primary
	}
	`
	instance := `{
		"Cars": [
			{ 	"regNbr": "ABC123",
		  		"MADE_BY_CarMakers": [{ "x": 1, "Where": { "name": "Fiat" } }, { "x": 2, "Where": { "name": "Skoda" } }]
			}
		],
		"CarMakers": [
			{ "name": "Fiat" },
			{ "name": "Skoda" }
		]
	}`
	inst := makeJzonInstance(t.Name(), instance)

	ctx, ic := makeContext(t.Name(), model, true)
	tt.CheckFalse(ic.HasFatal() || ic.HasErrors())

	tl := &yammm.PluggableListener{
		FOnEdge: func(ctx yammm.Context, a *yammm.Association, assocPropmap map[string]any,
			fromType *yammm.Type, fromPks map[string]any, toType *yammm.Type, toPks map[string]any,
		) {
			if i, ok := assocPropmap["x"].(int64); ok {
				switch i {
				case 1:
					tt.CheckEqual("Car", fromType.Name)
					tt.CheckEqual("CarMaker", toType.Name)
					tt.CheckEqual("ABC123", fromPks["regNbr"])
					tt.CheckEqual("Fiat", toPks["name"])
				case 2:
					tt.CheckEqual("Car", fromType.Name)
					tt.CheckEqual("CarMaker", toType.Name)
					tt.CheckEqual("ABC123", fromPks["regNbr"])
					tt.CheckEqual("Skoda", toPks["name"])
				default:
					t.Fatalf("Test got neither call 1 nor call 2")
				}
			}
		},
	}
	walker := yammm.NewGraphWalker(ctx, tl)
	walker.Walk(inst)
}
func Test_ListenerEnterExitComposition(t *testing.T) {
	tt := testutils.NewTester(t)

	model := `schema "testing"
	type Car {
		*-> HAS (one) Engine
		*-> HAS (many) AutoPart
	}
	part type Engine {
		power Integer
		-->MADE_BY (one) Maker
	}
	part type AutoPart {
		name String
	}
	type Maker {
		name String primary
	}
	`
	instance := `{
		"Cars": [
			{ 	"id": "$$the_car",
		  		"HAS_Engine": { "id": "$$1", "power": 100, "MADE_BY": {"Where": {"name": "RR"}} },
		  		"HAS_AutoParts": [ {"id": "$$2", "name": "ABS"}, {"id": "$$3", "name": "AluWheels"} ]
			}
		],
		"Makers": [ {"id": "$$rr", "name": "RR"}]
	}`
	inst := makeJzonInstance(t.Name(), instance)

	ctx, ic := makeContext(t.Name(), model, true)
	tt.CheckFalse(ic.HasFatal() || ic.HasErrors())

	visitedInstances := []string{}
	visitedCompositionEdges := []string{}
	tl := &yammm.PluggableListener{
		FEnterComposition: func(ctx yammm.Context, t *yammm.Type, c *yammm.Composition, data xray.Wrapper) {
			tt.CheckEqual("HAS", c.Name)
		},
		FExitComposition: func(ctx yammm.Context, t *yammm.Type, c *yammm.Composition, data xray.Wrapper) {
			tt.CheckEqual("HAS", c.Name)
		},
		FEnterInstance: func(ctx yammm.Context, t *yammm.Type, data xray.Wrapper) {
			// Entered 4 times, collect an easy to check string representation
			visitedInstances = append(visitedInstances, fmt.Sprintf("%s.%s", t.Name, data.Value("id")))
		},
		FOnEdge: func(ctx yammm.Context, a *yammm.Association, assocPropMap map[string]any,
			fromType *yammm.Type, fromPks map[string]any, toType *yammm.Type, toPks map[string]any) {
			tt.CheckEqual("Engine", fromType.Name)
			tt.CheckEqual("Maker", toType.Name)
			tt.CheckEqual("$$1", fromPks["id"])
			tt.CheckEqual("RR", toPks["name"])
		},
		FOnCompositionEdge: func(ctx yammm.Context, c *yammm.Composition,
			fromType *yammm.Type, fromPks map[string]any, toType *yammm.Type, toPks map[string]any) {
			visitedCompositionEdges = append(visitedCompositionEdges,
				fmt.Sprintf("%s.%s(%s -> %s)", fromType.Name, c.PropertyName(ctx), fromPks["id"], toPks["id"]))
		},
	}
	walker := yammm.NewGraphWalker(ctx, tl)
	walker.Walk(inst)
	tt.CheckNotNil(utils.Find(visitedInstances, func(s string) bool { return s == "Car.$$the_car" }))
	tt.CheckNotNil(utils.Find(visitedInstances, func(s string) bool { return s == "Engine.$$1" }))
	tt.CheckNotNil(utils.Find(visitedInstances, func(s string) bool { return s == "AutoPart.$$2" }))
	tt.CheckNotNil(utils.Find(visitedInstances, func(s string) bool { return s == "AutoPart.$$3" }))
	tt.CheckNotNil(utils.Find(visitedInstances, func(s string) bool { return s == "Maker.$$rr" }))

	tt.CheckNotNil(utils.Find(visitedCompositionEdges,
		func(s string) bool { return s == "Car.HAS_Engine($$the_car -> $$1)" }))
	tt.CheckNotNil(utils.Find(visitedCompositionEdges,
		func(s string) bool { return s == "Car.HAS_AutoParts($$the_car -> $$2)" }))
	tt.CheckNotNil(utils.Find(visitedCompositionEdges,
		func(s string) bool { return s == "Car.HAS_AutoParts($$the_car -> $$3)" }))
}

// Test composition with association.
