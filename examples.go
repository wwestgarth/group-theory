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

type Zn struct {
	modulus int
}

func (z *Zn) Operate(a, b groups.Element) groups.Element {
	bValue, _ := b.(int) // proper handle if type assertion fails
	aValue, _ := a.(int)
	return (aValue + bValue) % z.modulus
}

func (z *Zn) Equals(a, b groups.Element) bool {
	aValue, aOk := a.(int)
	bValue, bOk := b.(int)
	return aOk && bOk && aValue == bValue
}

func generatedZNGroup(zN int) {

	generated := groups.NewGroupFromGenerator(&Zn{zN}, 1, 500)
	result, err := generated.Validate()
	if err != nil {
		fmt.Println(err)
		return
	}

	printResults(fmt.Sprintf("Z%d", zN), result)
	return
}

func main() {

	generatedZNGroup(5)
	generatedZNGroup(6)
	generatedZNGroup(257) // prime group
	generatedZNGroup(360)

}
