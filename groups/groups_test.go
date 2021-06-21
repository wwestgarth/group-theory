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

func TestZ4GroupGiven(t *testing.T) {

	var g = createZnGroup(4)

	g.elements[0] = true
	g.elements[1] = true
	g.elements[2] = true
	g.elements[3] = true

	err := g.Analyse()
	if err != nil {
		t.Errorf("failed: %s", err)
	}
	g.FindGenerators()

	if !g.equals(g.identity, 0) {
		t.Errorf("Z5 identity = %d; want 0", g.identity)
	}

	if len(g.generators) != 2 {
		t.Errorf("Expected onlt 2 group generator, found %d", len(g.generators))
	}

	if len(g.elements) != 4 {
		t.Errorf("Z5 Order = %d; want 5", len(g.elements))
	}

	if !g.HasSubgroups() {
		t.Errorf("Z4 should have subgroups")
	}

}

func TestZ5GroupGenerated(t *testing.T) {

	order := 5
	var g = createZnGroup(order)

	if err := g.Generate(1, 6); err != nil {
		t.Errorf("Z5 could not generate group from %d", 1)
	}

	g.Analyse()

	if !g.equals(g.identity, 0) {
		t.Errorf("Z5 identity = %d; want 0", g.identity)
	}

	if len(g.generators) != 1 {
		t.Errorf("Could not find generators for group")
	}

	if len(g.elements) != order {
		t.Errorf("Z5 Order = %d; want 5", len(g.elements))
	}

	if g.HasSubgroups() {
		t.Errorf("Z5 should not have subgroups")
	}
}

func TestZ313GroupGenerated(t *testing.T) {

	order := 313
	var g = createZnGroup(order)

	if err := g.Generate(1, 315); err != nil {
		t.Errorf("Z313 could not generate group from %d", 1)
	}

	g.Analyse()

	if !g.equals(g.identity, 0) {
		t.Errorf("Z313 identity = %d; want 0", g.identity)
	}

	if len(g.elements) != order {
		t.Errorf("Z313 Order = %d; want 313", len(g.elements))
	}

	if g.HasSubgroups() {
		t.Errorf("Z313 should not have subgroups")
	}
}

func TestZLargeGroupGenerated(t *testing.T) {

	var g = createZnGroup(1000)
	if err := g.Generate(1, 6); err != ErrNotClosed {
		t.Errorf("Somehow managed to generate a group bigger that maxOrder")
	}
}

func TestGroupNotClosed(t *testing.T) {

	var g = createZnGroup(5)

	g.elements[1] = true
	g.elements[2] = true
	g.elements[3] = true
	g.elements[4] = true
	g.elements[5] = true

	var err = g.Analyse()

	if err == nil {
		t.Errorf("Expected to be unable to find identity")
	}
	if err != ErrNotClosed {
		t.Errorf("Wrong error code expected: %s got %s", ErrNotClosed, err)
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

	g.elements[1] = true
	g.elements[2] = true
	g.elements[3] = true
	g.elements[4] = true
	g.elements[5] = true

	var err = g.Analyse()

	if err == nil {
		t.Errorf("Found identity unexpectedly: %d", g.identity)
	} else if err != ErrNotClosed {
		t.Errorf("Wrong error code expected: %s got %s", ErrNotClosed, err)
	}
}
