package groups

// TODO: Fix isgenerator check
// TODO: Confirm whether identity check is correct
// TODO: Fix that a generator is non-unique
// TODO: add function comments
// Error handle, with proper error return types

import (
	"fmt"
)

type Element interface {
}

type GroupOperation func(Element, Element) Element
type GroupEquals func(Element, Element) bool

type Group struct {
	elements    map[Element]bool
	operator    GroupOperation
	equals      GroupEquals
	generator   Element
	cayleytable map[Element]map[Element]bool
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

func (g *Group) isidentity(e Element) bool {
	return g.equals(e, g.Operate(e, e)) // Unsure if this is sufficient
}

func (g *Group) isgenerator(e Element) bool {

	if g.generator != nil {
		return g.equals(g.generator, e) // Already cached
	}

	if g.isidentity(e) {
		return false // element is the identity
	}

	var current = e
	// Not quite right, this is either give the generator of the group
	// Or a generator of a sub-group. A generator is also not unique
	for i := 0; i < len(g.elements); i++ {
		current = g.Operate(e, current)
	}

	return g.equals(e, current)
}

func (g *Group) ensuregenerator() {

	if g.generator != nil {
		return
	}

	for element := range g.elements {
		if g.isgenerator(element) {
			g.generator = element
			return
		}
	}

}

func (g *Group) Analyse() {

	fmt.Println("Closed:", g.check_closure())

	g.ensuregenerator()

}

func (g *Group) Details() {
	fmt.Println("Group Details")
	fmt.Println("Order    :", len(g.elements))
	fmt.Println("Generator:", g.generator)
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
