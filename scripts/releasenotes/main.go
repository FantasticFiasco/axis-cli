package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	releaseNotes, err := readFromFile("./CHANGELOG.md")
	if err != nil {
		panic(err)
	}

	fmt.Print(releaseNotes)
}

func readFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	return readFrom(file)
}

func readFrom(src io.Reader) (string, error) {
	buf := new(strings.Builder)

	_, err := io.Copy(buf, src)
	if err != nil {
		return "", err
	}

	//return buf.String(), nil
	return "", nil
}
