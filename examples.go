package main

import (
	"fmt"
	"reflect"

	"github.com/wwestgarth/group-theory/groups"
)

func goperation(a, b groups.Element) groups.Element {

	aval := int(reflect.ValueOf(a).Int())
	bval := int(reflect.ValueOf(b).Int())
	var res int
	res = (aval + bval) % 5
	return res
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

	generated := groups.NewGroupFromGenerator(&group_op, &group_eq, 1, 10)
	generated.Analyse()
	generated.Details()

	explicit := groups.NewGroup(&group_op, &group_eq, set)
	explicit.Details()

	err := explicit.Analyse()
	explicit.Details()
	if err != nil {
		fmt.Println(err.Error())
	}
}
