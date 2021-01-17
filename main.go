package main

import (
	"fmt"
	"io"
)

type myReader struct {
	s   string
	inf int
	sup int
}

func (r *myReader) Read(b []byte) (n int, err error) {
	var aux string
	r.sup = len(b) + r.inf
	if r.sup > len(r.s) && r.inf < len(r.s) {
		r.sup = len(r.s)
	}
	if r.sup > len(r.s) {
		return 0, io.EOF
	}
	aux = r.s[r.inf:r.sup]
	r.inf = r.sup
	n = len(aux)
	for i, v := range aux {
		b[i] = byte(v)
	}
	return
}

func main() {
	var r = &myReader{s: "Insert text here!"}
	buf := make([]byte, 4)
	for {
		n, err := r.Read(buf)
		fmt.Println(n, err, string(buf[:n]))
		if err == io.EOF {
			break
		}
	}
	fmt.Println(buf)
}
