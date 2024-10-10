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

	curr []byte
	err  error
}

func (iter *ReaderSeparatedIterator) resetBuf() {
	iter.buf = nil
	iter.similarN = 0
}

func (iter *ReaderSeparatedIterator) Next() bool {
	defer iter.resetBuf()

	if iter.err != nil {
		return false
	}

	for iter.readerIter.Next() {
		var val byte
		iter.err = iter.readerIter.Scan(&val)
		if iter.err != nil {
			return false
		}

		iter.buf = append(iter.buf, val)
		if val == iter.separator[iter.similarN] {
			iter.similarN++
		} else {
			iter.similarN = 0
		}

		if iter.similarN == len(iter.separator) {
			iter.curr = iter.buf[:len(iter.buf)-len(iter.separator)]
			return true
		}
	}

	if len(iter.buf) > 0 {
		iter.curr = iter.buf
		return true
	}

	iter.err = ErrStopIteration
	return false
}

func (iter *ReaderSeparatedIterator) Scan(b *[]byte) error {
	if b != nil {
		*b = iter.curr
	}
	return iter.err
}
