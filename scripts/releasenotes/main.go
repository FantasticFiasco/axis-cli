package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

	err = ioutil.WriteFile("./RELEASE_NOTES.md", []byte(releaseNotes), 0744)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
