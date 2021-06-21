package main

import (
	"fmt"

	"github.com/wwestgarth/group-theory/groups"
)

func printResults(name string, result *groups.GroupValidateResult) {

	fmt.Println("Group name:    ", name)
	fmt.Println("Order:         ", result.Order)
	fmt.Println("Identity:      ", result.Identity)
	fmt.Println("Has subgroups: ", result.HasSubgroups)
	fmt.Println("Generators:    ", result.Generators)
	fmt.Println()

}

func generatedZNGroup(zN int) {

	// Make our group operations
	gOp := groups.GroupOperation(
		func(a, b groups.Element) groups.Element {
			bValue, _ := b.(int) // proper handle if type assertion fails
			aValue, _ := a.(int)
			return (aValue + bValue) % zN
		})

	gEq := groups.GroupEquals(
		func(a, b groups.Element) bool {
			aValue, aOk := a.(int)
			bValue, bOk := b.(int)
			return aOk && bOk && aValue == bValue
		})

	generated := groups.NewGroupFromGenerator(&gOp, &gEq, 1, 10)
	result, err := generated.Validate()
	if err != nil {
		return
	}

	printResults(fmt.Sprintf("Z%d", zN), result)
	return
}

func main() {

	generatedZNGroup(6)
	generatedZNGroup(5)

}
