package gosuper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceIterator_Next(t *testing.T) {
	originalSlice := []int{1, 2, 3}
	iter := SliceIterator[int]{slice: originalSlice}

	for i, expected := range originalSlice {
		actual, err := iter.Next()

		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
		assert.Len(t, iter.slice, len(originalSlice)-i-1)
		assert.Len(t, originalSlice, 3)
	}

	_, err := iter.Next()
	assert.ErrorIs(t, err, ErrStopIteration)
}

func TestSliceIterator_Next_Empty(t *testing.T) {
	iter := SliceIterator[int]{slice: []int{}}

	_, err := iter.Next()
	assert.ErrorIs(t, err, ErrStopIteration)
}
