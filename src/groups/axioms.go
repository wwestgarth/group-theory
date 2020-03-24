package groups

// checkClosure checks whether the elements in the Group acted on
// by the Group's Operator is closed 
func (g *Group) checkClosure() (bool, error) {

	var err GroupError
	g.cayleytable = make(map[Element]map[Element]bool)

	for element1 := range g.elements {
		for element2 := range g.elements {

			res := g.Operate(element1, element2)

			if !g.elements[res] {
				err.New(ErrorNotClosed)
				return false, err
			}

			_, is_in := g.cayleytable[element1]
			if !is_in {
				g.cayleytable[element1] = make(map[Element]bool)
			}

			g.cayleytable[element1][element2] = true
		}
	}

	return true, err
}
