package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(buf []byte) (n int, err error) {
	n, err = rot.r.Read(buf)

	for in, c := range buf {
		if (c >= 'a' && c <= 'm') || (c >= 'A' && c <= 'M') {
			buf[in] = c + 13
		} else if (c >= 'n' && c <= 'z') || (c >= 'N' && c <= 'Z') {
			buf[in] = c - 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
