package gosuper

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewReaderSeparatedIterator(t *testing.T) {
	originalReader := strings.NewReader("a,b,c")
	iter := NewReaderSeparatedIterator(originalReader, []byte(","))

	for _, expected := range []string{"a", "b", "c"} {
		actual, err := iter.Next()

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		assert.Equal(t, expected, string(actual))
	}

	_, err := iter.Next()
	assert.ErrorIs(t, err, ErrStopIteration)
}
