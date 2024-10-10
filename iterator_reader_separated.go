package gosuper

import (
	"errors"
	"io"
)

var ErrEmptySeparator = errors.New("empty separator")

func NewReaderSeparatedIterator(reader io.Reader, separator []byte) *ReaderSeparatedIterator {
	return &ReaderSeparatedIterator{readerIter: ReaderIterator{reader: reader}, separator: separator}
}

type ReaderSeparatedIterator struct {
	readerIter ReaderIterator
	separator  []byte
	buf        []byte
	similarN   int
}

func (iter *ReaderSeparatedIterator) resetBuf() {
	iter.buf = nil
	iter.similarN = 0
}

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
