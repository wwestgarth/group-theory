package main

import (
	"groups"
	"reflect"
)

func goperation(a, b groups.Element) groups.Element {

	aval := reflect.ValueOf(a).Int()
	bval := reflect.ValueOf(b).Int()
	return (aval + bval) % 5
}

func gequals(a, b groups.Element) bool {

	aval := reflect.ValueOf(a).Int()
	bval := reflect.ValueOf(b).Int()
	return aval == bval
}

func main() {

	var set = []groups.Element{0, 1, 2, 3, 4}

	var group_eq groups.GroupEquals
	group_eq = gequals

	var group_op groups.GroupOperation
	group_op = goperation

	generated := groups.New(&group_op, &group_eq)
	explicit := groups.New(&group_op, &group_eq)
	
	generated.Generate(1, 10)
	generated.Details()

	explicit.Add(set)
	explicit.Details()

	explicit.Analyse()
	explicit.Details()
}
