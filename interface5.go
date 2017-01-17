package main

import (
	"bytes"
	"io"
	"os"
)

type Reader struct {
	data []byte
}

func (r *Reader) eof() bool {
	return len(r.data) == 0
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.eof() {
		return 0, io.EOF
	}

	p[0] = r.data[0]
	r.data = r.data[1:]
	n = 1
	return
}

func main() {
	io.Copy(os.Stdout, bytes.NewReader([]byte("Hello\n")))

	r := &Reader{[]byte("Hello World\n")}
	io.Copy(os.Stdout, r)

}
