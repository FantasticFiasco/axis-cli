package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFrom(t *testing.T) {
	testCases := []struct {
		changelog    string
		version      string
		releaseNotes string
	}{
		{"testdata/1.0.0_CHANGELOG.md", "1.0.0", "testdata/1.0.0_RELEASENOTES.md"},
		{"testdata/1.0.0_CHANGELOG.md", "v1.0.0", "testdata/1.0.0_RELEASENOTES.md"},
		{"testdata/1.0.1_CHANGELOG.md", "1.0.1", "testdata/1.0.1_RELEASENOTES.md"},
		{"testdata/1.0.1_CHANGELOG.md", "v1.0.1", "testdata/1.0.1_RELEASENOTES.md"},
		{"testdata/1.1.0_CHANGELOG.md", "1.1.0", "testdata/1.1.0_RELEASENOTES.md"},
		{"testdata/1.1.0_CHANGELOG.md", "v1.1.0", "testdata/1.1.0_RELEASENOTES.md"},
		{"testdata/1.1.0_UNRELEASED_CHANGELOG.md", "v1.1.0", "testdata/1.1.0_UNRELEASED_RELEASENOTES.md"},
		{"testdata/1.1.0_UNRELEASED_CHANGELOG.md", "v1.1.0", "testdata/1.1.0_UNRELEASED_RELEASENOTES.md"},
	}

	for _, testCase := range testCases {
		want := mustReadTestData(testCase.releaseNotes)
		got, err := readFromFile(testCase.changelog, testCase.version)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, want, got)
	}
}

func mustReadTestData(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(content)
}
