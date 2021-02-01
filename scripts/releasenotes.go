package releasenotes

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	releaseNotes, err := ReadFromFile("./CHANGELOG.md")
	if err != nil {
		panic(err)
	}

	fmt.Print(releaseNotes)
}

// ReadFromFile returns the release notes from a change log file.
func ReadFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	return ReadFrom(file)
}

// ReadFrom returns the release notes from a change log reader.
func ReadFrom(src io.Reader) (string, error) {
	buf := new(strings.Builder)

	_, err := io.Copy(buf, src)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
