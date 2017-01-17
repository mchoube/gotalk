package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	io.Copy(os.Stdout, bytes.NewReader([]byte("Hello\n")))

}
