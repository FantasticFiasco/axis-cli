package main

import (
	"flag"
	"fmt"

	"github.com/FantasticFiasco/axis-cli/internal/build"
)

func main() {
	version := flag.Bool("version", false, "Show axis version")
	flag.Parse()

	if *version == true {
		fmt.Printf("axis version %s (%s)\n", build.Version, build.Date)
	}
}
