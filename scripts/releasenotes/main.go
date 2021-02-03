package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
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

	return findChapterContent(file, version)
}

func findChapterContent(r io.Reader, version string) (string, error) {
	startPrefix := fmt.Sprintf("## [%s]", version)
	stopPrefix := "## "

	isReadingContent := false
	content := ""

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), startPrefix) {
			isReadingContent = true
		} else if isReadingContent {
			if strings.HasPrefix(scanner.Text(), stopPrefix) {
				break
			} else {
				content += scanner.Text()
			}
		}
	}

	return content, scanner.Err()
}

func chapterRegexp() (*regexp.Regexp, error) {

	expr := "## " + // Markdown chapter
		`\[` + // Begin bracket
		`(` + // Begin capturing group
		`\d+\.\d+\.\d+` + // The version
		`)` + // End capturing group
		`\]` // End bracket

	return regexp.Compile(expr)
}
