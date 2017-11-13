package main

import (
	"io"
	"os"
	"strings"
)

// ROT13 为凯撒加密变体，字母对应为其13位之后的字母

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	_,  e := rot.r.Read(b)
	for k, v := range b {
		if 65<=v && 122>=v {
			if t := v+13; t <= 122 {
				b[k] = t
			} else {
				b[k] = 65 + t - 122
			}
		}
	}
	return len(b), e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}