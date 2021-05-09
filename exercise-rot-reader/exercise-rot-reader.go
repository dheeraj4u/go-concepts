package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rotFX(b byte) byte {
	uct := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lct := []byte("abcdefghijklmnopqrstuvwxyz")
	if pos := bytes.IndexByte(uct, b); pos != -1 {
		return uct[(pos+13)%len(uct)]
	}
	if pos := bytes.IndexByte(lct, b); pos != -1 {
		return lct[(pos+13)%len(lct)]
	}
	return b
}
func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rotFX(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
