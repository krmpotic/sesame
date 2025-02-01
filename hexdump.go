package main

import (
	"fmt"
	"io"
)

func hexdump(w io.Writer, p []byte) {
	for i := range p {
		if i%16 == 0 {
			if i != 0 {
				fmt.Fprint(w, "\n")
			}
			fmt.Fprintf(w, "%07x ", i)
		} else if i%2 == 0 {
			fmt.Fprintf(w, " ")
		}

		fmt.Fprintf(w, "%0x", p[i])
	}
	if len(p)%16 != 0 {
		fmt.Fprintf(w, "\n")
	}
	fmt.Fprintf(w, "%07x\n", len(p))
}
