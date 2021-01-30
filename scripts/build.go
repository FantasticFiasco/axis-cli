// Build tasks for the project

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		usage(os.Stderr)
		os.Exit(1)
	}
}

func usage(w io.Writer) {
	fmt.Fprintln(w, "Usage: go run scripts/build.go <task>")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Known tasks are:")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  bin/axis:")
	fmt.Fprintln(w, "    Builds the main executable.")
	fmt.Fprintln(w, "    Supported environment variables:")
	fmt.Fprintln(w, "    - GO_LDFLAGS")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  clean:")
	fmt.Fprintln(w, "    Deletes all built files.")
	fmt.Fprintln(w, "")
}
