package gosuper

// NewSliceIterator creates a new SliceIterator for the given slice.
func NewSliceIterator[Type any](slice []Type) *SliceIterator[Type] {
	return &SliceIterator[Type]{slice: slice}
}

// SliceIterator is an iterator implementation for slices.
type SliceIterator[Type any] struct {
	slice []Type
}

// Next returns the next element in the slice.
// If there are no more elements, it returns the zero value for Type and ErrStopIteration.
func (iter *SliceIterator[Type]) Next() (t Type, err error) {
	if len(iter.slice) == 0 {
		return t, ErrStopIteration
	}

	next := iter.slice[0]
	iter.slice = iter.slice[1:]
	return next, nil
}
