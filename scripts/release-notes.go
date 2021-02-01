package main

import (
	"fmt"
	"os"
)

func main() {
	w := os.Stdout

	fmt.Fprintln(w, "# Hello world")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "## Chapter 1")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Some text")
}
