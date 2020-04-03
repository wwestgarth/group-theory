package groups

// groupIsClosed checks whether the elements in the Group acted on
// by the Group's Operator is closed
func groupIsClosed(g *Group) (bool, error) {

	var err GroupError
	g.cayleytable = make(map[Element]map[Element]bool)

	for element1 := range g.elements {
		for element2 := range g.elements {

			res := g.Operate(element1, element2)

			if !g.elements[res] {
				err.New(ErrorNotClosed, element2, element2)
				return false, err
			}

			_, is_in := g.cayleytable[element1]
			if !is_in {
				g.cayleytable[element1] = make(map[Element]bool)
			}

			g.cayleytable[element1][element2] = true
		}
	}

	return true, nil
}

// groupHasIdentity Checks and returns the identity element of the suspected group
func groupHasIdentity(g *Group) (Element, error) {

	var err GroupError
	var identity Element

	for element1 := range g.elements {
		for element2 := range g.elements {

			res := g.Operate(element1, element2)

			if g.equals(element2, res) {
				identity = element1
			}
		}
	}

	if identity == nil {
		err.New(ErrorNoIdentity, nil, nil)
		return nil, err
	}
	return identity, nil
}

// findInverseElement Given and Element finds its inverse in the Group
func findInverseElement(g *Group, e Element) (Element, error) {

	var err GroupError
	var inverse Element

	if g.identity == nil {
		err.New(ErrorNoIdentity, nil, nil)
	}

	if g.equals(g.identity, e) {
		return e, nil
	}

	for element1 := range g.elements {

		res := g.Operate(e, element1)

		if g.equals(res, g.identity) {
			inverse = element1
			break
		}
	}

	if inverse == nil {
		err.New(ErrorNoInverse, e, nil)
		return nil, err
	}

	return inverse, nil
}

// groupHasInverses Returns true if ever element in thr group has an inverse
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
