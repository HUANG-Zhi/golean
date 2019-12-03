package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(b []byte) (int,error){
	i,e := r13.r.Read(b)
	if i > 0 {
		l,h := byte('a'), byte('z')
		L,H := byte('A'), byte('Z')
		mod := h - l + 1
		for j:=0; j < i; j++ {
			if b[j] >= l && b[j] <= h {
				b[j] = (b[j] + 13 - l) % mod + l
			}else if b[j] >= L && b[j] <= H {
				b[j] = (b[j] + 13 - L) % mod + L
			}
		}
	}
	return i,e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

