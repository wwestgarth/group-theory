package groups

import (
	"errors"
	"fmt"

	"github.com/wwestgarth/group-theory/primality"
)

var (

	// ErrorOpFailed generic specified error
	ErrGroupOpFailed = errors.New("group Operation failed")
	// ErrorNotClosed Groups is not closed
	ErrNotClosed = errors.New("group is not closed")
	// ErrorNoIdentity Group does not have an Identity
	ErrNoIdentity = errors.New("group does not have an Identity")
	// ErrorNoInverse Group element does not have an inverse
	ErrNoInverse = errors.New("group element does not have an inverse")
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
	elements   map[Element]bool
	operator   GroupOperation
	equals     GroupEquals
	generators []Element
	identity   Element
	table      *cayleyTable
}

// New returns a new instance of a Group. Requires an Groups operation
// and a means of element equality.
func New(op *GroupOperation, eq *GroupEquals) Group {
	var g Group
	g.elements = make(map[Element]bool)
	g.operator = *op
	g.equals = *eq
	g.table = newCayleyTable()
	return g
}

// New returns a new instance of a Group. Requires an Groups operation
// and a means of element equality.
func NewGroup(op *GroupOperation, eq *GroupEquals, elements []Element) (g Group) {

	g.operator = *op
	g.equals = *eq

	g.elements = make(map[Element]bool)
	g.table = newCayleyTable()

	// Fill in elements
	for _, element := range elements {
		g.elements[element] = true
	}

	return g
}

// New returns a new instance of a Group. Requires an Groups operation
// and a means of element equality.
func NewGroupFromGenerator(op *GroupOperation, eq *GroupEquals, generator Element, maxOrder uint32) (g Group) {
	g = New(op, eq)
	g.Generate(generator, maxOrder)

	return
}

// Operate executes the Group's registered operator, using the Cayley Table
// to look up values if the operation has been performed before
func (g *Group) Operate(a, b Element) (value Element) {

	value, err := g.table.lookup(a, b)
	if err != nil {
		value = g.operator(a, b)
	}

	return
}

// Generate Attempts to generates a group from the given generator. If more
// than 'max_order' elements are added to the Group then we group generate
// is stopped.
func (g *Group) Generate(generator Element, maxOrder uint32) (err error) {

	var current = generator

	found := false
	for i := uint32(0); i < maxOrder; i++ {

		g.elements[current] = true
		current = g.Operate(generator, current)

		if g.equals(generator, current) {
			found = true
			break
		}
	}

	if !found {
		err = ErrNotClosed
		return
	}

	return
}

// isGenerator returns true if the given Element is a generator of the Group
func (g *Group) isGenerator(e Element) (generator bool) {

	if g.isIdentity(e) {
		return // element is the identity, can't generate group
	}

	current := e
	for i := 0; i < len(g.elements)-1; i++ {

		current = g.Operate(e, current)
		if g.equals(e, current) {
			return // generator of sub-group
		}
	}

	current = g.Operate(e, current)
	generator = g.equals(e, current)
	return
}

// FindGenerators find generators in the group
func (g *Group) FindGenerators() {

	for element := range g.elements {
		if g.isGenerator(element) {
			g.generators = append(g.generators, element)
		}
	}
}

// HasSubgroups return whether the group has subgroups based on the primatlity of the group order
func (g *Group) HasSubgroups() bool {
	return !primality.ProbablyPrime(len(g.elements), 10)
}

// Validate checks the given Group's elements and Operation satisies the axioms
// of Group Theory
func (g *Group) Validate() (result *GroupValidateResult, err error) {

	// Axiom 1: Closure
	if err = g.isClosed(); err != nil {
		return
	}

	// Axiom 2: Existence of identity
	if _, err = g.findIdentity(); err != nil {
		return
	}

	// Axiom 3: All elements have an inverse
	if err = g.hasInverses(); err != nil {
		return
	}

	// Axiom 4: Associativity
	// Not worth checking...?

	g.FindGenerators()
	// Fill in things we know
	result = &GroupValidateResult{
		Order:        uint32(len(g.elements)),
		Identity:     g.identity,
		HasSubgroups: g.HasSubgroups(),
		Generators:   g.generators[:],
	}
	return
}

type GroupValidateResult struct {
	Order        uint32
	Identity     Element
	Generators   []Element
	HasSubgroups bool
}

// Details prints all know details of the Group
func (g *Group) Details() {
	fmt.Println("Group Details")
	fmt.Println("Order     :", len(g.elements))
	fmt.Println("Generators:", g.generators)
	fmt.Println("Identity  :", g.identity)
}
