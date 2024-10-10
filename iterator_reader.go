package gosuper

import "io"

func NewReaderIterator(reader io.Reader) *ReaderIterator {
	return &ReaderIterator{reader: reader}
}

type ReaderIterator struct {
	reader io.Reader
	buf    []byte
	curr   byte
	err    error
}

func (iter *ReaderIterator) Next() bool {
	if iter.reader == nil {
		return false
	}

	if len(iter.buf) == 0 {
		iter.buf = make([]byte, 1)
	}

	_, iter.err = iter.reader.Read(iter.buf)
	if iter.err != nil {
		if iter.err == io.EOF {
			iter.reader = nil
			iter.err = ErrStopIteration
			return false
		}
	}

	iter.curr = iter.buf[0]
	return true
}

func (iter *ReaderIterator) Scan(b *byte) error {
	if b != nil {
		*b = iter.curr
	}
	return iter.err
}
