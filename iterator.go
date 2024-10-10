package gosuper

import "errors"

// ErrStopIteration is returned when an iterator has no more elements.
var ErrStopIteration = errors.New("stop iteration")

// Iterator is a generic interface for iterating over a collection of elements.
type Iterator[Type any] interface {
	// Next returns the next element in the iteration.
	// If there are no more elements, it returns the zero value for Type and ErrStopIteration.
	Next() bool
	Scan(*Type) error
}

// Iterable is an interface for types that can produce an Iterator.
type Iterable[Type any] interface {
	// Iterator returns an Iterator over the elements of the collection.
	Iterator() Iterator[Type]
}
