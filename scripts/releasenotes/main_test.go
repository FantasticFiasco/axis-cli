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
		{"testdata/BARE_CHANGELOG.md", "v1.0.0", "testdata/BARE_RELEASENOTES.md"},
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
