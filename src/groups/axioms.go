package groups

func (g *Group) check_closure() bool {

	g.cayleytable = make(map[Element]map[Element]bool)

	for element1 := range g.elements {
		for element2 := range g.elements {

			res := g.Operate(element1, element2)

			if !g.elements[res] {
				return false
			}

			_, is_in := g.cayleytable[element1]
			if !is_in {
				g.cayleytable[element1] = make(map[Element]bool)
			}

			g.cayleytable[element1][element2] = true
		}
	}

	return true
}
