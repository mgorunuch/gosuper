package gosuper

import "errors"

var ErrStopIteration = errors.New("stop iteration")

type Iterator[Type any] interface {
	Next() bool
	Scan(*Type) error
}

type Iterable[Type any] interface {
	Iterator() Iterator[Type]
}
