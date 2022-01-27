package main

import (
	"fmt"
	"io"
	"os"
)

// Capper implements io.Writer and turns everything to uppercase
type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A')
	out := make([]byte, len(p))

	fmt.Printf("diff value %v type %T\n", diff, diff)
	fmt.Printf("out type %T\n", out)

	for i, c := range p {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}
		out[i] = c
	}

	fmt.Println("in", p)
	fmt.Println("out", out)

	return c.wtr.Write(out)
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello there")
}
