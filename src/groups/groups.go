package groups

type Element interface {
}

type GroupOperation func(Element, Element) Element

type Group struct {
	elements map[Element]bool
	operator GroupOperation
}

func (g *Group) Add(elements []Element) {

	for _, element := range elements {
		g.elements[element] = true
	}
}

func (g *Group) RegisterOperation(operation *GroupOperation) {
	g.operator = *operation
}

func (g *Group) Operate(a, b Element) Element {
	return g.operator(a,b)
}

func New() Group {
	var g Group
	g.elements = make(map[Element]bool)
	return g
}

