package groups

import "errors"

var (
	ErrCayleyEntryNotFound = errors.New("entry not in table")
)

type cayleyTable struct {
	table map[Element]map[Element]Element
}

func newCayleyTable() (c *cayleyTable) {
	c = new(cayleyTable)
	c.table = make(map[Element]map[Element]Element)
	return
}

func (c *cayleyTable) add(a, b, res Element) {

	if _, ok := c.table[a]; !ok {
		c.table[a] = make(map[Element]Element)
	}

	c.table[a][b] = res
	return
}

func (c *cayleyTable) lookup(a, b Element) (value Element, err error) {

	err = ErrCayleyEntryNotFound
	if _, ok := c.table[a]; !ok {
		return
	}

	value, ok := c.table[a][b]

	if !ok {
		return
	}

	err = nil
	return
}
