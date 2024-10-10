package gosuper

func NewSliceIterator[Type any](slice []Type) *SliceIterator[Type] {
	return &SliceIterator[Type]{slice: slice}
}

type SliceIterator[Type any] struct {
	slice []Type
}

func (iter *SliceIterator[Type]) Next() (t Type, err error) {
	if len(iter.slice) == 0 {
		return t, ErrStopIteration
	}

	next := iter.slice[0]
	iter.slice = iter.slice[1:]
	return next, nil
}
