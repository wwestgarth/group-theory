package groups

// groupIsClosed checks whether the elements in the Group acted on
// by the Group's Operator is closed
func groupIsClosed(g *Group) (closed bool) {

	for element1 := range g.elements {
		for element2 := range g.elements {

			res := g.Operate(element1, element2)

			if !g.elements[res] {
				closed = false
				return
			}

			if _, ok := g.cayleytable[element1]; !ok {
				g.cayleytable[element1] = make(map[Element]Element)
			}

			g.cayleytable[element1][element2] = res
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
	return g.equals(e, g.Operate(e, e))
}

// groupHasIdentity Checks and returns the identity element of the suspected group
func groupHasIdentity(g *Group) (identity Element, err error) {

	for element := range g.elements {
		if g.isIdentity(element) {
			identity = element
			return
		}
	}

	return nil, ErrNoIdentity
}

// findInverseElement Given and Element finds its inverse in the Group
func findInverseElement(g *Group, e Element) (inverse Element, err error) {

	if g.identity == nil {
		err = ErrNoIdentity
		return
	}

	// given element is identity
	if g.equals(g.identity, e) {
		inverse = e
		return
	}

	for element1 := range g.elements {

		res := g.Operate(e, element1)

		if g.equals(res, g.identity) {
			inverse = element1
			return
		}
	}

	err = ErrNoInverse
	return
}

// groupHasInverses Returns true if every element in thr group has an inverse
func groupHasInverses(g *Group) (res bool, err error) {

	if g.identity == nil {
		err = ErrNoIdentity
		return
	}

	for element1 := range g.elements {

		if _, err = findInverseElement(g, element1); err != nil {
			return
		}
	}

	res = true
	return
}
