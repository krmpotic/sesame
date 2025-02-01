package main

import (
	"errors"
	"io"
)

var ErrNoSpace = errors.New("no space")

type NibbleWriter struct {
	p []byte
}

func NewNibbleWriter(p []byte) *NibbleWriter {
	return &NibbleWriter{p}
}

func (w *NibbleWriter) Write(p []byte) (int, error) {
	if len(w.p) < 2 {
		return 0, ErrNoSpace
	}

	if 2*len(p) > len(w.p) {
		p = p[:len(w.p)/2]
	}

	for _, b := range p {
		w.writeByte(b)
	}

	return len(p), nil
}

func (w *NibbleWriter) writeByte(b byte) {
	n0 := byte(0xf0 & b >> 4)
	w.p[0] &= 0xf0
	w.p[0] |= n0
	n1 := byte(0x0f & b)
	w.p[1] &= 0xf0
	w.p[1] |= n1
	w.p = w.p[2:]
}

type NibbleReader struct {
	p []byte
}

func NewNibbleReader(p []byte) *NibbleReader {
	return &NibbleReader{p}
}

func (r *NibbleReader) Read(p []byte) (int, error) {
	if len(r.p) < 2 {
		return 0, io.EOF
	}
	if len(p) > 2*len(r.p) {
		p = p[:len(r.p)/2]
	}
	for i := range p {
		p[i] = r.readByte()
	}
	return len(p), nil
}

func (r *NibbleReader) readByte() (b byte) {
	b = r.p[0]<<4 | r.p[1]&0x0f
	r.p = r.p[2:]
	return b
}
