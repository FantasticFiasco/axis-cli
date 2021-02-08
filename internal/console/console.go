package console

import (
	"fmt"
	"os"
)

func Print(a ...interface{}) {
	fmt.Print(a...)
}

func Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func Fatal(a ...interface{}) {
	fmt.Fprint(os.Stderr, a...)
	os.Exit(1)
}

func Fatalf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}
