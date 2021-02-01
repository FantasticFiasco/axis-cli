package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	version := flag.String("version", "", "The version to get release notes for")
	flag.Parse()

	if *version == "" {
		flag.Usage()
		os.Exit(1)
	}

	releaseNotes, err := readFromFile("./CHANGELOG.md", *version)
	if err != nil {
		panic(err)
	}

	fmt.Print(releaseNotes)
}

func readFromFile(filename, version string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return readFrom(file, version)
}

func readFrom(src io.Reader, version string) (string, error) {
	buf := new(strings.Builder)

	_, err := io.Copy(buf, src)
	if err != nil {
		return "", err
	}

	//return buf.String(), nil
	return "", nil
}
