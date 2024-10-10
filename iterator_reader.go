package gosuper

import "io"

// NewReaderIterator creates a new ReaderIterator for the given io.Reader.
func NewReaderIterator(reader io.Reader) *ReaderIterator {
	return &ReaderIterator{reader: reader}
}

// ReaderIterator is an iterator implementation for io.Reader.
type ReaderIterator struct {
	reader io.Reader
	buf    []byte
}

// Next reads and returns the next byte from the reader.
// If there are no more bytes to read, it returns 0 and ErrStopIteration.
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
