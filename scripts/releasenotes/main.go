package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	version := flag.String("version", "", "The version to get release notes for")
	flag.Parse()

	if *version == "" {
		flag.Usage()
		os.Exit(1)
	}

	releaseNotes, err := ReadFromFile("./CHANGELOG.md", *version)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(releaseNotes)
}
