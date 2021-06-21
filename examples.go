package main

import (
	"fmt"

	"github.com/wwestgarth/group-theory/groups"
)

func goperation(a, b groups.Element) groups.Element {

	aValue, ok := a.(int)
	if !ok {
		// Need to handle this
		return false
	}

	bValue, ok := b.(int)
	if !ok {
		// Need to handle this
		return false
	}

	return (aValue + bValue) % 6
}

func gequals(a, b groups.Element) bool {

	aValue, ok := a.(int)
	if !ok {
		// Need to handle this
		return false
	}

	bValue, ok := b.(int)
	if !ok {
		// Need to handle this
		return false
	}

	return aValue == bValue
}

func detailsGeneratedGroup() {
	return
}

func main() {

	var set = []groups.Element{0, 1, 2, 3, 4, 5}

	var group_eq groups.GroupEquals
	group_eq = gequals

	var group_op groups.GroupOperation
	group_op = goperation

	generated := groups.NewGroupFromGenerator(&group_op, &group_eq, 1, 10)
	generated.Validate()
	generated.Details()

	explicit := groups.NewGroup(&group_op, &group_eq, set)
	explicit.Details()

	_, err := explicit.Validate()
	explicit.Details()
	if err != nil {
		fmt.Println(err.Error())
	}
}
