package go_super

import (
	"strings"
	"testing"
)

func TestReaderIterator_Next(t *testing.T) {
	originalReader := strings.NewReader("abc")
	iter := ReaderIterator{reader: originalReader}

	for _, expected := range []byte{'a', 'b', 'c'} {
		actual, err := iter.Next()

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if expected != actual {
			t.Errorf("expected %v, got %v", expected, actual)
		}
		if iter.reader != originalReader {
			t.Errorf("unexpected reader change")
		}
	}

	_, err := iter.Next()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestReaderIterator_Next_Empty(t *testing.T) {
	iter := ReaderIterator{reader: strings.NewReader("")}

	_, err := iter.Next()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
