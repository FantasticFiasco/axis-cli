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
	version = strings.TrimPrefix(version, "v")
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
				content += scanner.Text() + "\n"
			}
		}
	}

	if scanner.Err() != nil {
		return "", scanner.Err()
	}

	content = strings.Trim(content, "\n")

	if content == "" {
		return "", fmt.Errorf("release with version %s was not found", version)
	}

	return content, nil
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
