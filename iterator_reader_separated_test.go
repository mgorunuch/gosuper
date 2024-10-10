package gosuper

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewReaderSeparatedIterator(t *testing.T) {
	originalReader := strings.NewReader("a,b,c")
	iter := NewReaderSeparatedIterator(originalReader, []byte(","))

	for _, expected := range []byte{'a', 'b', 'c'} {
		assert.True(t, iter.Next())

		var actual []byte
		err := iter.Scan(&actual)

		assert.NoError(t, err)
		assert.Equal(t, []byte{expected}, actual)
	}

	assert.False(t, iter.Next())
	assert.ErrorIs(t, iter.Scan(nil), ErrStopIteration)
}
