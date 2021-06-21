package groups

import (
	"errors"
)

var (

	// ErrorOpFailed generic specified error
	ErrGroupOpFailed = errors.New("Group Operation failed")
	// ErrorNotClosed Groups is not closed
	ErrNotClosed = errors.New("Group is not closed")
	// ErrorNoIdentity Group does not have an Identity
	ErrNoIdentity = errors.New("Group does not have an Identity")
	// ErrorNoInverse Group element does not have an inverse
	ErrNoInverse = errors.New("Group element does not have an inverse")
)
