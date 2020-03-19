package groups

import (
	"fmt"
)

type Element interface {
}

type GroupOperation func(Element, Element) Element
type GroupEquals func(Element, Element) bool

type Group struct {
	elements map[Element]bool
	operator GroupOperation
	equals GroupEquals
	generator Element
}

func (g *Group) Add(elements []Element) {

	for _, element := range elements {
		g.elements[element] = true
	}
}

func (g *Group) Generate(generator Element, max_order int) bool {

	found := false
	var current = generator
	var elements []Element

	for i := 0; i < max_order; i++ {

		elements = append(elements, current)
		current = g.Operate(generator, current)

		if g.equals(generator, current) {
			found = true
			break
		}
	}

	if found {
		g.generator = generator
		g.Add(elements)
	}

	return found
}

func (g *Group) Operate(a, b Element) Element {
	return g.operator(a, b)
}

func (g *Group) Details() {
	fmt.Println("Group Details");
	fmt.Println("Order    :", len(g.elements));
	fmt.Println("Generator:", g.generator);
}

func New(op *GroupOperation, eq *GroupEquals) Group {
	var g Group
	g.elements = make(map[Element]bool)
	g.operator = *op
	g.equals = *eq
	return g
}
