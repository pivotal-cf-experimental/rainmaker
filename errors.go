package rainmaker

import "github.com/pivotal-cf-experimental/rainmaker/internal/network"

type NotFoundError struct {
	err error
}

func (e NotFoundError) Error() string {
	return e.err.Error()
}

type Error struct {
	err error
}

func (e Error) Error() string {
	return e.err.Error()
}

func translateError(err error) error {
	switch err.(type) {
	case network.NotFoundError:
		return NotFoundError{err}
	default:
		return Error{err}
	}
}
