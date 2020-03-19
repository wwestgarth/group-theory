package groups

// TODO: Fix isgenerator check
// TODO: add function comments

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

func (g *Group) isgenerator(e Element) bool{

	if g.generator != nil {
		return g.equals(g.generator, e) // Already cached
	}

	if g.equals(e, g.Operate(e, e)) {
		return false // element is the identity...or its own inverse
	}

	var current = e
	// Not quite right, this is either give the generator of the group
	// Or a generator of a sub-group.
	for i := 0; i < len(g.elements); i++ {
		current = g.Operate(e, current)
		if
	}
	
	return g.equals(e, current)
}

func (g *Group) ensuregenerator() {

	if g.generator != nil {
		fmt.Println("Early")
		fmt.Println(g.generator)
		return
	}
	
	for element := range g.elements {
		if g.isgenerator(element) {
			g.generator = element
			return
		}
	}

}

func (g *Group) Analyse()  {

	g.ensuregenerator()

}

func (g *Group) Details() {
	fmt.Println("Group Details");
	fmt.Println("Order    :", len(g.elements));
	fmt.Println("Generator:", g.generator);
}

func (g *Group) Operate(a, b Element) Element {
	return g.operator(a, b)
}

func New(op *GroupOperation, eq *GroupEquals) Group {
	var g Group
	g.elements = make(map[Element]bool)
	g.operator = *op
	g.equals = *eq
	return g
}
