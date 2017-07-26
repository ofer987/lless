package main

import (
	"fmt"
	"io"
)

const Version = "0.1.0"

func displayVersion(w io.Writer) {
	fmt.Fprintf(w, "lless v%s\n", Version)
}
