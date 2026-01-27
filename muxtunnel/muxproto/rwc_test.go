package muxproto

import (
	"io"
	"testing"
)

type payload struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Random string `json:"random"`
}

func TestAuth(t *testing.T) {
	src := new(stupidWC)
	p := new(payload)
	WriteAuth(src, p)

	dr := newDelimReader(src)
	dest := new(stupidWC)
	io.CopyBuffer(dest, dr, make([]byte, 1))

	t.Logf("%s\n", dest)
}

func newDelimReader(r io.Reader) *delimReader {
	return &delimReader{raw: r, dem: authEOF}
}

type stupidWC struct {
	buf []byte
}

func (sr *stupidWC) Write(p []byte) (int, error) {
	sr.buf = append(sr.buf, p...)
	return len(p), nil
}

func (sr *stupidWC) Read(p []byte) (int, error) {
	if len(sr.buf) == 0 {
		return 0, io.EOF
	}

	n := copy(p, sr.buf)
	sr.buf = sr.buf[n:]

	return n, nil
}
