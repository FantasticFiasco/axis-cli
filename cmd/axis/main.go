package main

import (
	"flag"
	"fmt"
)

// The following variables are set by goreleaser during CD
var version = "<version>"
var commit = "<git sha>"
var date = "<date>"

func main() {
	versionFlag := flag.Bool("version", false, "Show axis version")
	flag.Parse()

	if *versionFlag == true {
		fmt.Printf("axis %s\n", version)
		fmt.Printf("commit:  %s\n", commit)
		fmt.Printf("release: %s\n", releaseURL())
		fmt.Printf("date:    %s\n", date)
	}
}

func releaseURL() string {
	url := "https://github.com/FantasticFiasco/axis-cli/releases"
	if version != "<version>" {
		url += "/tag/" + version
	}
	return url
}
