package main

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFrom(t *testing.T) {
	testCases := []struct {
		changelog    string
		releaseNotes string
	}{
		{"BARE_CHANGELOG.md", "BARE_RELEASENOTES.md"},
	}

	for _, testCase := range testCases {
		want := mustReadTestData(testCase.releaseNotes)
		got, err := readFromFile(path.Join("testdata", testCase.changelog))
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, want, got)
	}
}

func mustReadTestData(filename string) string {
	content, err := ioutil.ReadFile(path.Join("testdata", filename))
	if err != nil {
		panic(err)
	}
	return string(content)
}
