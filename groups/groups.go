package groups

import (
	"github.com/wwestgarth/group-theory/primality"
)

// generate Attempts to generates a group from the given generator. If more
// than 'max_order' elements are added to the Group then we group generate
// is stopped.
func (g *Group) generate(generator Element, maxOrder uint32) (err error) {

	var current = generator

	found := false
	for i := uint32(0); i < maxOrder; i++ {

		g.elements[current] = true
		current = g.Operate(generator, current)

		if g.Equals(generator, current) {
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
		if g.Equals(e, current) {
			return // generator of sub-group
		}
	}

	current = g.Operate(e, current)
	generator = g.Equals(e, current)
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
