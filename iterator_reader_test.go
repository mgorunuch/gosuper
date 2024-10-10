package gosuper

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestReaderIterator_Next(t *testing.T) {
	originalReader := strings.NewReader("abc")
	iter := ReaderIterator{reader: originalReader}

	for _, expected := range []byte{'a', 'b', 'c'} {
		assert.True(t, iter.Next())

		var actual byte
		assert.NoError(t, iter.Scan(&actual))
		assert.Equal(t, expected, actual)
	}

	assert.False(t, iter.Next())
	assert.ErrorIs(t, iter.Scan(nil), ErrStopIteration)
}

func TestReaderIterator_Next_Empty(t *testing.T) {
	iter := ReaderIterator{reader: strings.NewReader("")}
	assert.False(t, iter.Next())
}
