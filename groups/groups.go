package groups

// TODO: Fix isgenerator check
// TODO: Confirm whether identity check is correct
// TODO: Fix that a generator is non-unique
// Error handle, with proper error return types

import (
	"fmt"
)

// Element An abstracted group element
type Element interface {
}

// GroupOperation An abtracted group operation defined on two group elements
type GroupOperation func(Element, Element) Element

// GroupEquals An abstracted equality functions defined on two group elements
type GroupEquals func(Element, Element) bool

// Group Structure which defines a Group
type Group struct {
	elements    map[Element]bool
	operator    GroupOperation
	equals      GroupEquals
	generator   Element
	identity    Element
	cayleytable map[Element]map[Element]bool
}

// Add adds the given slice to the Group. These are the members of the Group.
func (g *Group) Add(elements []Element) {

	for _, element := range elements {
		g.elements[element] = true
	}
}

// Generate Attempts to generates a group from the given generator. If more
// than 'max_order' elements are added to the Group then we group generate
// is stopped.
func (g *Group) Generate(generator Element, maxOrder int) bool {

	found := false
	var current = generator
	var elements []Element

	for i := 0; i < maxOrder; i++ {

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

// isIdentity returns true if the given Element is the Group's Identity
func (g *Group) isIdentity(e Element) bool {
	return g.equals(e, g.Operate(e, e)) // Unsure if this is sufficient
}

// isGenerator returns true if the given Element is a geneartor of the Group
func (g *Group) isGenerator(e Element) bool {

	if g.generator != nil {
		return g.equals(g.generator, e) // Already cached
	}

	if g.isIdentity(e) {
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

// ensureGenerator ensures the Group's generator field is known
func (g *Group) ensureGenerator() {

	if g.generator != nil {
		return
	}

	for element := range g.elements {
		if g.isGenerator(element) {
			g.generator = element
			return
		}
	}
}

// Analyse checks the given Group's elements and Operation satisies the axioms
// of Group Theory
func (g *Group) Analyse() error {

	_, err := groupIsClosed(g)
	if err != nil {
		return err
	}

	g.identity, err = groupHasIdentity(g)
	if err != nil {
		return err
	}

	_, err = groupHasInverses(g)
	if err != nil {
		return err
	}

	g.ensureGenerator()

	return nil
}

// Details prints all know details of the Group
func (g *Group) Details() {
	fmt.Println("Group Details")
	fmt.Println("Order    :", len(g.elements))
	fmt.Println("Generator:", g.generator)
	fmt.Println("Identity :", g.identity)
}

// Operate executes the Group's registered operator
func (g *Group) Operate(a, b Element) Element {
	return g.operator(a, b)
}

// New returns a new instance of a Group. Requires an Groups operation
// and a means of element equality.
func New(op *GroupOperation, eq *GroupEquals) Group {
	var g Group
	g.elements = make(map[Element]bool)
	g.operator = *op
	g.equals = *eq
	return g
}
