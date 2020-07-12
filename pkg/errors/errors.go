package errors

import "github.com/pkg/errors"

type Kind uint8

const (
	Other         Kind = iota // Unclassified error..
	RequestFailed             // Request is failed.
	NotFound                  // Resource does not exists.
)

const (
	CodeNotFound      = "not-found"
	CodeAlreadyExists = "already-exist"
	CodeInternalError = "internal-error"
)

type Error struct {
	Kind Kind
	Code string
	Err  error
}

func (e *Error) Error() string {
	return e.Err.Error()
}

// E builds an error value from its arguments.c
func E(kind Kind, code string, str string) error {
	return &Error{
		Kind: kind,
		Code: code,
		Err:  errors.New(str),
	}
}

// Is reports whether err is an *Error of the given Kind.
func Is(kind Kind, err error) bool {
	e, ok := err.(*Error)
	if !ok {
		return false
	}
	return e.Kind == kind
}
