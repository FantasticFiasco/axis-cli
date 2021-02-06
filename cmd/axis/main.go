package main

import (
	"flag"
	"github.com/FantasticFiasco/axis-cli/cmd/axis/commands"
)

// The following variables are set by goreleaser during CD
var version = "<version>"
var commit = "<git sha>"
var date = "<date>"

func main() {
	versionFlag := flag.Bool("version", false, "Show axis version")
	flag.Parse()

	if *versionFlag == true {
		commands.Version(version, commit, date)
	} else {
		flag.Usage()
	}
}
