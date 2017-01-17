package main

import (
	"bytes"
	"io"
	"os"
)

type Reader struct {
	data []byte
}

func main() {
	io.Copy(os.Stdout, bytes.NewReader([]byte("Hello\n")))

	r := Reader{[]byte("Hello World\n")}
	io.Copy(os.Stdout, r)

}
