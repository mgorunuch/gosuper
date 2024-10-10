package gosuper

import (
	"errors"
	"io"
)

// ErrEmptySeparator is returned when an empty separator is provided.
var ErrEmptySeparator = errors.New("empty separator")

// NewReaderSeparatedIterator creates a new ReaderSeparatedIterator for the given io.Reader and separator.
func NewReaderSeparatedIterator(reader io.Reader, separator []byte) *ReaderSeparatedIterator {
	return &ReaderSeparatedIterator{readerIter: ReaderIterator{reader: reader}, separator: separator}
}

// ReaderSeparatedIterator is an iterator implementation that separates input from an io.Reader based on a given separator.
type ReaderSeparatedIterator struct {
	readerIter ReaderIterator
	separator  []byte
	buf        []byte
	similarN   int
}

// resetBuf resets the internal buffer and similarity counter.
func (iter *ReaderSeparatedIterator) resetBuf() {
	iter.buf = nil
	iter.similarN = 0
}

// Next returns the next chunk of bytes up to the separator.
// If there are no more chunks to read, it returns nil and ErrStopIteration.
func (iter *ReaderSeparatedIterator) Next() ([]byte, error) {
	defer iter.resetBuf()

	for {
		val, err := iter.readerIter.Next()
		if err != nil {
			if errors.Is(err, ErrStopIteration) {
				if len(iter.buf) > 0 {
					return iter.buf, nil
				}
			}

			return nil, err
		}

		if len(iter.separator) == 0 {
			return nil, ErrEmptySeparator
		}

		iter.buf = append(iter.buf, val)
		if val == iter.separator[iter.similarN] {
			iter.similarN++
		} else {
			iter.similarN = 0
		}

		if iter.similarN == len(iter.separator) {
			return iter.buf[:len(iter.buf)-len(iter.separator)], nil
		}
	}
}
