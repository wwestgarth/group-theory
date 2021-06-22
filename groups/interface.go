package groups

import "errors"

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

type Operator interface {
	Operate(a, b Element) Element
	Equals(a, b Element) bool
}

// Group Structure which defines a Group
type Group struct {
	elements   map[Element]bool
	op         Operator
	generators []Element
	identity   Element
	table      *cayleyTable
}

// GroupValidateResults gives insight into the nature of the given valid group
type GroupValidateResult struct {
	Order        uint32
	Identity     Element
	Generators   []Element
	HasSubgroups bool
}

// initGroup initialises all the insides of a group and seeds the NewGroupXXX functions
func initGroup(op Operator) Group {
	var g Group
	g.elements = make(map[Element]bool)
	g.table = newCayleyTable()
	g.op = op
	return g
}

// New returns a new instance of a Group. Requires an Groups operation
// and a means of element equality.
func NewGroup(op Operator, elements []Element) (g Group) {

	g = initGroup(op)

	// Fill in elements
	for _, element := range elements {
		g.elements[element] = true
	}

	return g
}

// New returns a new instance of a Group. Requires an Groups operation
// and a means of element equality.
func NewGroupFromGenerator(op Operator, generator Element, maxOrder uint32) (g Group) {
	g = initGroup(op)
	g.generate(generator, maxOrder)

	return
}

// Operate executes the Group's registered operator, using the Cayley Table
// to look up values if the operation has been performed before
func (g *Group) Operate(a, b Element) (value Element) {

	value, err := g.table.lookup(a, b)
	if err != nil {
		value = g.op.Operate(a, b)
	}

	return
}

// Operate executes the Group's registered operator, using the Cayley Table
// to look up values if the operation has been performed before
func (g *Group) Equals(a, b Element) bool {
	return g.op.Equals(a, b)
}
