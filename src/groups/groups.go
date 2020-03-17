package groups

import (
	"fmt"
)


type Element interface {
}

type GroupOperation func(Element, Element) Element

type Group struct {
	elements map[Element]bool
	operator GroupOperation
}

func (g Group) Add(elements []Element) {

	for _, element := range elements {
		fmt.Println("Adding", element)
		g.elements[element] = true
	}
}

func (g Group) RegisterOperation(operation *GroupOperation) {
	g.operator = *operation
	fmt.Println(g.operator)

}

func (g Group) Operate(a, b Element) Element {
	fmt.Println("Operate")
	fmt.Println(g.operator)
	return g.operator(a,b)
}

func New() Group {
	var g Group
	g.elements = make(map[Element]bool)
	return g
}

