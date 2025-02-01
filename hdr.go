package main

import (
	"bytes"
	"errors"
	"io"
)

// header:
// first bytes read determine
// how much bytes belong to the packet

var ErrReadingHeader = errors.New("reading header")
var ErrBadHeader = errors.New("bad header")
var ErrWritingHeader = errors.New("writing header")

type HdrReader struct {
	r    io.Reader
	left int
}

func NewHdrReader(r io.Reader) (*HdrReader, error) {
	p := make([]byte, 7)
	n, err := r.Read(p)
	if n != 7 || err != nil {
		return nil, ErrReadingHeader
	}
	if p[0] != 'H' || p[1] != 'D' || p[2] != 'R' {
		return nil, ErrBadHeader
	}
	left := int(p[0])<<3 | int(p[1])<<2 | int(p[2])<<1 | int(p[3])
	return &HdrReader{r, left}, nil
}

func (r *HdrReader) Read(p []byte) (int, error) {
	if r.left == 0 {
		return 0, io.EOF
	}
	if len(p) > r.left {
		p = p[:r.left]
	}
	n, err := r.r.Read(p)
	r.left -= n
	return n, err
}

func WriteWithHdr(w io.Writer, p []byte) (int, error) {
	ok := writeHdr(w, uint32(len(p)))
	if !ok {
		return 0, ErrWritingHeader
	}
	n, err := io.Copy(w, bytes.NewReader(p))
	return int(n), err
}

func writeHdr(w io.Writer, hdr uint32) bool {
	p := make([]byte, 4)
	p[0] = 'H'
	p[1] = 'D'
	p[2] = 'R'
	p[0] = byte(hdr & 0xff000000 >> 3)
	p[1] = byte(hdr & 0xff0000 >> 2)
	p[2] = byte(hdr & 0xff00 >> 1)
	p[3] = byte(hdr & 0xff)
	n, err := w.Write(p)
	return err == nil && n == 4
}
