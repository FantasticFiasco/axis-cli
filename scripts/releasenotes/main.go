package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	versionFlag := flag.String("version", "", "The version to get release notes for")
	flag.Parse()

	if *versionFlag == "" {
		flag.Usage()
		os.Exit(1)
	}

	releaseNotes, err := ReadFromFile("./CHANGELOG.md", *versionFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(releaseNotes)
}
