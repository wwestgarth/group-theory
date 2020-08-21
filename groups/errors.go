package groups

const (
	ErrorOpFailed   = "Group Operation failed"
	ErrorNotClosed  = "Group is not closed"
	ErrorNoIdentity = "Group does not have an Identity"
	ErrorNoInverse  = "Group element does not have an inverse"
)

type GroupError struct {
	Code string
	Err1 Element
	Err2 Element
}

func (e *GroupError) New(code string, e1, e2 Element) {
	e.Code = code
	e.Err1 = e1
	e.Err1 = e2
}

func (e GroupError) Error() string { return e.Code }
