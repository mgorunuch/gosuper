package gosuper

import "io"

func NewReaderIterator(reader io.Reader) *ReaderIterator {
	return &ReaderIterator{reader: reader}
}

type ReaderIterator struct {
	reader io.Reader
	buf    []byte
}

func (iter *ReaderIterator) Next() (t byte, err error) {
	if iter.reader == nil {
		return t, ErrStopIteration
	}

	if len(iter.buf) == 0 {
		iter.buf = make([]byte, 1)
	}

	_, err = iter.reader.Read(iter.buf)
	if err != nil {
		if err == io.EOF {
			iter.reader = nil
			return t, ErrStopIteration
		}

		return t, err
	}

	return iter.buf[0], nil
}
