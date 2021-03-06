package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func ReadFromFile(filename, version string) (string, error) {
	matched, err := regexp.MatchString(`^\d+.\d+.\d+$`, version)
	if err != nil {
		return "", err
	}
	if !matched {
		return "", fmt.Errorf(`"%s" does not conform to the expected version format`, version)
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return readChapterContent(file, version)
}

func readChapterContent(r io.Reader, version string) (string, error) {
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
