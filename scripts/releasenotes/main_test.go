package main_test

import (
	"github.com/FantasticFiasco/axis-cli/scripts/releasenotes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFromFile(t *testing.T) {
	testCases := []struct {
		changelog    string
		version      string
		releaseNotes string
	}{
		{"testdata/1.0.0_CHANGELOG.md", "1.0.0", "testdata/1.0.0_RELEASENOTES.md"},
		{"testdata/1.0.1_CHANGELOG.md", "1.0.1", "testdata/1.0.1_RELEASENOTES.md"},
		{"testdata/1.1.0_CHANGELOG.md", "1.1.0", "testdata/1.1.0_RELEASENOTES.md"},
		{"testdata/1.1.0_UNRELEASED_CHANGELOG.md", "1.1.0", "testdata/1.1.0_UNRELEASED_RELEASENOTES.md"},
	}

	for _, testCase := range testCases {
		want := mustReadTestData(testCase.releaseNotes)
		got, err := main.ReadFromFile(testCase.changelog, testCase.version)
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	}
}

func TestReadFromGivenInvalidVersion(t *testing.T) {
	_, err := main.ReadFromFile("testdata/1.0.0_CHANGELOG.md", "2.0.0")
	assert.NotNil(t, err)
}

func mustReadTestData(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Trim(string(content), "\r\n")
}
