package groups

// groupIsClosed checks whether the elements in the Group acted on
// by the Group's Operator is closed
func groupIsClosed(g *Group) (bool, error) {

	var err GroupError

	for element1 := range g.elements {
		for element2 := range g.elements {

			res := g.Operate(element1, element2)

			if !g.elements[res] {
				err.New(ErrorNotClosed, element2, element2)
				return false, err
			}

			_, isIn := g.cayleytable[element1]
			if !isIn {
				g.cayleytable[element1] = make(map[Element]Element)
			}

			g.cayleytable[element1][element2] = res
		}
	}

	return true, nil
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
func groupHasIdentity(g *Group) (Element, error) {

	var err GroupError
	var identity Element

	for element := range g.elements {
		if g.isIdentity(element) {
			identity = element
			return identity, nil
		}
	}

	err.New(ErrorNoIdentity, nil, nil)
	return nil, err
}

// findInverseElement Given and Element finds its inverse in the Group
func findInverseElement(g *Group, e Element) (Element, error) {

	var err GroupError

	if g.identity == nil {
		err.New(ErrorNoIdentity, nil, nil)
		return nil, err
	}

	if g.equals(g.identity, e) {
		return e, nil
	}

	for element1 := range g.elements {

		res := g.Operate(e, element1)

		if g.equals(res, g.identity) {
			return element1, nil
		}
	}

	err.New(ErrorNoInverse, e, nil)
	return nil, err
}

// groupHasInverses Returns true if every element in thr group has an inverse
func groupHasInverses(g *Group) (bool, error) {

	var err GroupError

	if g.identity == nil {
		err.New(ErrorNoIdentity, nil, nil)
		return false, err
	}

	for element1 := range g.elements {

		_, err := findInverseElement(g, element1)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
