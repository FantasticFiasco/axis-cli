// Build tasks for the project

package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		args = append(args, "usage")
	}

	var err error

	switch args[0] {
	case "bin/axis":
		break
	case "clean":
		err = rmrf("bin")
		break
	default:
		err = usage()
		break
	}

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}

func usage() error {
	return errors.New(
		strings.Join([]string{
			"Usage: go run scripts/build.go <task>",
			"",
			"Known tasks are:",
			"",
			"  bin/axis:",
			"    Builds the main executable.",
			"    Supported environment variables:",
			"    - GO_LDFLAGS",
			"",
			"  clean:",
			"    Deletes all built files.",
			"",
		},
			"\n"),
	)
}

func rmrf(targets ...string) error {
	args := append([]string{"rm", "-rf"}, targets...)
	print(args)
	for _, target := range targets {
		if err := os.RemoveAll(target); err != nil {
			return err
		}
	}
	return nil
}

func print(args []string) {
	fmt.Println(strings.Join(args, " "))
}
