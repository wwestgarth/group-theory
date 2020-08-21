package groups

const (
	// ErrorOpFailed generic specified error
	ErrorOpFailed = "Group Operation failed"
	// ErrorNotClosed Groups is not closed
	ErrorNotClosed = "Group is not closed"
	// ErrorNoIdentity Group does not have an Identity
	ErrorNoIdentity = "Group does not have an Identity"
	// ErrorNoInverse Group element does not have an inverse
	ErrorNoInverse = "Group element does not have an inverse"
)

// GroupError structure defining an error-type returned by the groups package
type GroupError struct {
	Code string
	Err1 Element
	Err2 Element
}

// New instantiates a new group error
func (e *GroupError) New(code string, e1, e2 Element) {
	e.Code = code
	e.Err1 = e1
	e.Err1 = e2
}

func (e GroupError) Error() string { return e.Code }
