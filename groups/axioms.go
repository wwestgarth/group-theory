package groups

// groupIsClosed checks whether the elements in the Group acted on
// by the Group's Operator is closed
func (g *Group) isClosed() (err error) {

	// This is a bit rough
	// Could improve performance if we assume communitive? But otherwise, by definition,
	// we have to check all pairs
	for element1 := range g.elements {
		for element2 := range g.elements {

			res := g.Operate(element1, element2)

			if !g.elements[res] {
				err = ErrNotClosed
				return
			}

			g.table.add(element1, element2, res)
		}
	}

	return
}

// isIdentity returns true if the given Element is the Group's Identity
// The identity of a group is unique, and has order 1. If the below holds
// true then order of e is either 1, or infinity. Given the order of an element
// must divide the group order, for finite groups the order cannot be infinity
// and so it must be 1, and thus the identity element.
func (g *Group) isIdentity(e Element) bool {
	return g.Equals(e, g.Operate(e, e))
}

// groupHasIdentity Checks and returns the identity element of the suspected group
func (g *Group) findIdentity() (identity Element, err error) {

	if g.identity != nil {
		identity = g.identity
		return
	}

	for element := range g.elements {

		if g.isIdentity(element) {
			identity = element
			g.identity = identity
			return
		}
	}

	err = ErrNoIdentity
	return
}

// findInverseElement Given and Element finds its inverse in the Group
func findInverseElement(g *Group, e Element) (inverse Element, err error) {

	if _, err = g.findIdentity(); err != nil {
		return
	}

	// given element is identity
	if g.Equals(g.identity, e) {
		inverse = e
		return
	}

	for element1 := range g.elements {

		res := g.Operate(e, element1)

		if g.Equals(res, g.identity) {
			inverse = element1
			return
		}
	}

	err = ErrNoInverse
	return
}

// groupHasInverses Returns true if every element in thr group has an inverse
func (g *Group) hasInverses() (err error) {

	if _, err = g.findIdentity(); err != nil {
		return
	}

	for element1 := range g.elements {

		if _, err = findInverseElement(g, element1); err != nil {
			return
		}
	}

	return
}
