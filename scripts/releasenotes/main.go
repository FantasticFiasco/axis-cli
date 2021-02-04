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

	return readChapterContent(file, version)
}

func readChapterContent(r io.Reader, version string) (string, error) {
	version = strings.TrimPrefix(version, "v")
	startPrefix := fmt.Sprintf("## [%s]", version)
	stopPrefix := "## "

	isWithinChapter := false
	chapterContent := ""

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), startPrefix) {
			isWithinChapter = true
		} else if isWithinChapter && strings.HasPrefix(scanner.Text(), stopPrefix) {
			break
		} else if isWithinChapter {
			chapterContent += scanner.Text() + "\n"
		}
	}

	if scanner.Err() != nil {
		return "", scanner.Err()
	}

	chapterContent = strings.Trim(chapterContent, "\n")
	if chapterContent == "" {
		return "", fmt.Errorf("release with version %s was not found in changelog", version)
	}

	return chapterContent, nil
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
