package groups

import (
	"reflect"
	"testing"
)

func gZoperation(modulus int) func(a, b Element) Element {

	return func(a, b Element) Element {
		aval := int(reflect.ValueOf(a).Int())
		bval := int(reflect.ValueOf(b).Int())
		var res int
		res = (aval + bval) % modulus
		return res
	}
}

func gZequals(a, b Element) bool {

	aval := reflect.ValueOf(a).Int()
	bval := reflect.ValueOf(b).Int()
	return aval == bval
}

func createZnGroup(modulus int) Group {
	var groupEq GroupEquals
	groupEq = gZequals

	var groupOp GroupOperation
	groupOp = gZoperation(modulus)
	g := New(&groupOp, &groupEq)

	return g
}

func TestZ5GroupGiven(t *testing.T) {

	var g = createZnGroup(5)
	var set = []Element{0, 1, 2, 3, 4}

	g.Add(set)

	g.Analyse()
	if !g.equals(g.identity, 0) {
		t.Errorf("Z5 identity = %d; want 0", g.identity)
	}

	if !g.equals(g.generator, 1) {
		t.Errorf("Z5 generator = %d; want 1", g.generator)
	}

	if len(g.elements) != 5 {
		t.Errorf("Z5 Order = %d; want 5", len(g.elements))
	}
}

func TestZ5GroupGenerated(t *testing.T) {

	var g = createZnGroup(5)

	if !g.Generate(1, 6) {
		t.Errorf("Z5 could not generate group from %d", 1)
	}

	g.Analyse()

	if !g.equals(g.identity, 0) {
		t.Errorf("Z5 identity = %d; want 0", g.identity)
	}

	if !g.equals(g.generator, 1) {
		t.Errorf("Z5 generator = %d; want 1", g.generator)
	}

	if len(g.elements) != 5 {
		t.Errorf("Z5 Order = %d; want 5", len(g.elements))
	}
}

func TestZLargeGroupGenerated(t *testing.T) {

	var g = createZnGroup(1000)
	if g.Generate(1, 6) {
		t.Errorf("Somehow managed to generate a group bigger that maxOrder")
	}
}

func TestGroupNotClosed(t *testing.T) {

	var g = createZnGroup(5)
	var set = []Element{1, 2, 3, 4, 5}
	g.Add(set)

	var err = g.Analyse()

	if err == nil {
		t.Errorf("Expected to be unable to find identity")
	}
	if err.Error() != ErrorNotClosed {
		t.Errorf("Wrong error code expected: %s got %s", ErrorNotClosed, err.Error())
	}
}

func noIdentityOp(a, b Element) Element {
	aval := int(reflect.ValueOf(a).Int())
	return aval + 1
}
func TestNoIdentity(t *testing.T) {

	var groupEq GroupEquals
	groupEq = gZequals

	var groupOp GroupOperation
	groupOp = noIdentityOp
	g := New(&groupOp, &groupEq)

	var set = []Element{1, 2, 3, 4, 5}
	g.Add(set)

	var err = g.Analyse()

	if err == nil {
		t.Errorf("Found identity unexpectedly: %d", g.identity)
	} else if err.Error() != ErrorNotClosed {
		t.Errorf("Wrong error code expected: %s got %s", ErrorNotClosed, err.Error())
	}
}
