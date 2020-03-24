package groups

const (
	ErrorOpFailed   = "Group Operation failed"
	ErrorNotClosed  = "Group is not closed"
	ErrorNoIdentity = "Group does not have an Identity"
)

type GroupError struct {
	Code string
	Err1 Element
	Err2 Element
}

func (e *GroupError) New(code string) {
	e.Code = code
}

func (e GroupError) Error() string { return e.Code }
