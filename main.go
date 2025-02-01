package main

import (
	"fmt"
	"io"
	"os"
)

func write(w io.Writer, r io.Reader, secret []byte) error {
	return nil
}

func read(r io.Reader) ([]byte, error) {
	return nil, nil
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	switch os.Args[1] {
	case "save":
		if len(os.Args) != 4 {
			usage()
		}

		secret, err := os.ReadFile(os.Args[2])
		if err != nil {
			fail("opening %v: %v", os.Args[2], err)
		}

		file, err := os.Open(os.Args[3])
		if err != nil {
			fail("opening %v: %v", os.Args[3], err)
		}
		defer file.Close()

		write(os.Stdout, file, secret)
	case "open":
		if len(os.Args) != 3 {
			usage()
		}

		file, err := os.Open(os.Args[2])
		if err != nil {
			fail("opening file: %v", err)
		}
		defer file.Close()

		secret, err := read(file)

		fmt.Println(secret)
	default:
		usage()
	}
}

func fail(str string, v ...any) {
	str = "error: " + str + "\n"
	fmt.Fprintf(os.Stderr, str, v...)
	os.Exit(1)
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: \n")
	fmt.Fprintf(os.Stderr, "	sesame open <file>\n")
	fmt.Fprintf(os.Stderr, "	sesame save <secret_file> <file>\n")
	os.Exit(1)
}
