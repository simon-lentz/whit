package yammm_test

import (
	"fmt"
	"testing"
)

func Test_Validator_Invariant(t *testing.T) {
	model := `schema "testing"
	type Subject {
		a Integer
		! "this is always true"
		  true
	}`
	instance := `{ "Subjects": [ { "a": 42 } ] }`
	validateMessages(t, model, instance) // expect no messages
}
func TestInvariant_Arithmetic(t *testing.T) {
	testOkInvariant(t, "add sub mul precedence", `1 + 2 * 3 + 2 -1 == 8`)
	testOkInvariant(t, "add sub div precedence", `1 + 8 / 2 - 2 == 3`)
	testOkInvariant(t, "modulo precedence", `20 - 10 % 3  == 19`)
}

func TestInvariant_BooleanOperators(t *testing.T) {
	testOkInvariant(t, "true", `true`)
	testOkInvariant(t, "not true", `!true == false`)
	testOkInvariant(t, "not not nil", `!!_ == false`) //nolint:dupword
	testOkInvariant(t, "not nil", `!_ == true`)

	testOkInvariant(t, "or", `false || true == true`)
	testOkInvariant(t, "or", `false || false == false`)

	testOkInvariant(t, "and", `true && true == true`)
	testOkInvariant(t, "and", `true && false == false`)

	testOkInvariant(t, "xor", `false ^ true == true`)
	testOkInvariant(t, "xor", `true ^ false == true`)
	testOkInvariant(t, "xor", `false ^ false == false`)
	testOkInvariant(t, "xor", `true ^ true == false`)
	testOkInvariant(t, "xor", `_ ^ true == true`)
	testOkInvariant(t, "xor", `true ^ _ == true`)

	testOkInvariant(t, "or nil", `true ^ true == false`)
}
func TestInvariant_If(t *testing.T) {
	testOkInvariant(t, "if true", `true ? { true : false} == true`)
	testOkInvariant(t, "if false", `false ? { false : true} == true`)
	testOkInvariant(t, "if _", `_ ? { false : true} == true`)
}
func TestInvariant_Equality(t *testing.T) {
	testOkInvariant(t, "== int", `1 == 1`)
	testOkInvariant(t, "== int", `1 == 2 == false`)
	testOkInvariant(t, "== float", `1.0 == 1.0`)
	testOkInvariant(t, "== float", `1.0 == 1.1 == false`)
	testOkInvariant(t, "== string", `"abc" == "abc"`)
	testOkInvariant(t, "== string", `"abx" == "abc" == false`)
	testOkInvariant(t, "== array", `[1,2,3] == [1,2,3]`)
	testOkInvariant(t, "== array", `[1,[2,3]] == [1,[2,3]]`)
	testOkInvariant(t, "== array", `[1,2,3] == [1,2,4] == false`)
	testOkInvariant(t, "== array", `[1,2,3] == [1,2,3,5] == false`)

	testOkInvariant(t, "== nil", `_ == _`)
	testOkInvariant(t, "!= nil", `_ == true == false`)
	testOkInvariant(t, "!= nil", `_ == false == false`)

	testOkInvariant(t, "!= int", `1 != 1 == false`)
	testOkInvariant(t, "!= int", `1 != 2`)
	testOkInvariant(t, "!= float", `1.0 != 1.0 == false`)
	testOkInvariant(t, "!= float", `1.0 != 1.1`)
	testOkInvariant(t, "!= string", `"abc" != "abc" == false`)
	testOkInvariant(t, "!= string", `"abx" != "abc"`)
	testOkInvariant(t, "!= array", `[1,2,3] != [1,2,3] == false`)
	testOkInvariant(t, "!= array", `[1,[2,3]] != [1,[2,3]] == false`)
	testOkInvariant(t, "!= array", `[1,2,3] != [1,2,4]`)
	testOkInvariant(t, "!= array", `[1,2,3] != [1,2,3,5]`)

	// TODO: nil, bool, regexp.
}
func TestInvariant_Comparisons(t *testing.T) {
	testOkInvariant(t, "< int", `1 < 2`)
	testOkInvariant(t, "< int", `1 < 0 == false`)
	testOkInvariant(t, "<= int", `1 <= 2`)
	testOkInvariant(t, "<= int", `2 <= 2`)
	testOkInvariant(t, "<= int", `2 <= 3 == true`)

	testOkInvariant(t, "< float", `1.0 < 2.0`)
	testOkInvariant(t, "< float", `1.0 < 0.0 == false`)
	testOkInvariant(t, "<= float", `1.0 <= 2.0`)
	testOkInvariant(t, "<= float", `2.0 <= 2.0`)
	testOkInvariant(t, "<= float", `2.0 <= 2.1 == true`)

	// TODO: string, array, nil, bool, regexp.
}
func TestInvariant_Match(t *testing.T) {
	testOkInvariant(t, "=~ int", `1 =~ Integer`)
	testOkInvariant(t, "=~ int", `1 =~ Integer[1,10]`)
	testOkInvariant(t, "=~ int", `1 =~ Integer[5,10] == false`)
	testOkInvariant(t, "=~ float", `1.2 =~ Float[1.0,10.0]`)
	testOkInvariant(t, "=~ float", `1.2 =~ Float[1,10]`)

	testOkInvariant(t, "!~ int", `1 !~ Integer == false`)
	testOkInvariant(t, "!~ int", `1 !~ Integer[1,10] == false`)
	testOkInvariant(t, "!~ int", `1 !~ Integer[5,10]`)
	testOkInvariant(t, "!~ float", `1.2 !~ Float[1.0,10.0] == false`)

	testOkInvariant(t, "=~ regexp", `"hello" =~ /el/`)
	testOkInvariant(t, "=~ regexp", `"goodbye" =~ /el/ == false`)
}

func TestInvariant_In(t *testing.T) {
	testOkInvariant(t, "in", `3 in [1,2,2,3,2]`)
	testOkInvariant(t, "in", `4 in [1,2,2,3,2] == false`)
	testOkInvariant(t, "in", `[1,2] in [[1,2], 3]`)
	testOkInvariant(t, "in", `[1,2] in [[1,3], 1,2] == false`)
}

func TestInvariant_Slice(t *testing.T) {
	testOkInvariant(t, "slice", `[1,2,3,4][1] == 2`)
	testOkInvariant(t, "slice", `[1,2,3,4][1,2] == [2]`)
	testOkInvariant(t, "slice", `[1,2,3,4][1,3] == [2,3]`)
}

func TestInvariant_Functions(t *testing.T) {
	testOkInvariant(t, "count three", `3 == [1,2,2,3,2]->count {$0 == 2}`)
	testOkInvariant(t, "count three", `[1,2,2,3,2]->count {$0 == 2} == 3`)
	testOkInvariant(t, "count three", `2->count {$0 == 2} == 1`)

	testOkInvariant(t, "map double", `[1,2,3]->map {$0*2} == [2,4,6]`)
	testOkInvariant(t, "map double", `1->map {$0*2} == [2]`)

	testOkInvariant(t, "filter 2's", `[1,2,2,3]->filter {$0 != 2} == [1,3]`)
	testOkInvariant(t, "filter 2's", `1->filter {$0 != 2} == [1]`)

	testOkInvariant(t, "reduce by sum", `[1,2,3,4]->reduce(-1) {$0 + $1} == 9`)
	testOkInvariant(t, "reduce by sum", `4->reduce(-1) {$0 + $1} == 3`)

	testOkInvariant(t, "then", `2 -> then { $0*10} == 20`)

	testOkInvariant(t, "lest", `2 -> lest { 3 } == 2`)
	testOkInvariant(t, "lest", `_ -> lest { 3 } == 3`)

	testOkInvariant(t, "with", `(2+2) -> with { $0 < 5 && $0 > 3} == true`)

	testOkInvariant(t, "compact", `[1,_,2,_, 3,_]->compact == [1,2,3]`)
	testOkInvariant(t, "compact", `_->compact == []`)

	testOkInvariant(t, "len", `[1,2,3]->len == 3`)
	testOkInvariant(t, "len", `[]->len == 0`)
	testOkInvariant(t, "len", `"abc"->len == 3`)

	testOkInvariant(t, "unique", `[1,2,2]->unique == [1,2]`)
	testOkInvariant(t, "unique", `[1,[9,8],[9,8]]->unique == [1,[9,8]]`)
	testOkInvariant(t, "unique", `[]->unique == []`)
	testOkInvariant(t, "unique", `_->unique == [_]`)
	testOkInvariant(t, "unique", `1->unique == [1]`)

	testOkInvariant(t, "all", `1->all { $0 != _} == true`)
	testOkInvariant(t, "all", `[1,2,3]->all { $0 != _} == true`)
	testOkInvariant(t, "all", `[1,2,3]->all { $0 == _} == false`)

	testOkInvariant(t, "all_or_none", `1->all_or_none { $0 != _} == true`)
	testOkInvariant(t, "all_or_none", `[1,2,3]->all_or_none { $0 != _} == true`)
	testOkInvariant(t, "all_or_none", `[1,2,_]->all_or_none { $0 == _} == false`)

	testOkInvariant(t, "any", `1->any { $0 != _} == true`)
	testOkInvariant(t, "any", `[1,2,3]->any { $0 != _} == true`)
	testOkInvariant(t, "any", `[1,2,3]->any { $0 == _} == false`)

	testOkInvariant(t, "compare", `1->compare(0) == 1`)
	testOkInvariant(t, "compare", `1->compare(2) == -1`)
	testOkInvariant(t, "compare", `1->compare(1) == 0`)
	// other forms of value comparisons extensively tested in utils TypeOrder.
}
func TestInvariant_MatchRegexp(t *testing.T) {
	testOkInvariant(t, "match", `"Hi Fred"->match(/re/) == ["re"]`)
	testOkInvariant(t, "match", `"Hi Fred"->match(/(\w+)\s(\w+)/) == ["Hi Fred", "Hi", "Fred"]`)
}
func TestMathFunctions(t *testing.T) {
	testOkInvariant(t, "abs", `-1->abs == 1`)
	testOkInvariant(t, "floor", `1.2->floor == 1.0`)
	testOkInvariant(t, "ceil", `1.2->ceil == 2.0`)
	testOkInvariant(t, "min", `1->min(2) == 1`)
	testOkInvariant(t, "min", `[3,2,3]->min == 2`)
	testOkInvariant(t, "round", `23.5->round == 24.0`)
	testOkInvariant(t, "round", `-23.5->round == -24.0`)
}
func TestVarScope(t *testing.T) {
	testOkInvariant(t, "x in inner scope", `[1,2,3]->with|$x| { $x->reduce { $x }} == [1,2,3]`)
	testOkInvariant(t, "$0 in inner scope", `[1,2,3]->with { $0->with |$x| { $0 }  } == _`)
	testOkInvariant(t, "not $0 in inner scope", `[1,2,3]->with { $0->with |$x| { $x }  } == [1,2,3]`)
}

func TestInvariantPropertyAccess(t *testing.T) {
	testOkInvariant(t, "value of direct access to property a", `a == 42`)
	testOkInvariant(t, "value of property a in $self", `$self.a == 42`)
	testOkInvariant(t, "value of property string a in $self", `$self."a" == 42`)
	testOkInvariant(t, "value of property via other var", `$self->with {$0.a} == 42`)
}

// TODO: undefined var should be an error.
func testOkInvariant(t *testing.T, msg, expr string) {
	t.Helper()
	model := fmt.Sprintf(`schema "testing"
	type Subject {
		a Integer
		b Integer
		c Integer
		! "%s"
		  %s
	}`, msg, expr)
	instance := `{ "Subjects": [ { "a": 42 } ] }`
	validateMessages(t, model, instance) // expect no messages
}
