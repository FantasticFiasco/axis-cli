package releasenotes_test

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
		{"CHANGELOG_BARE.md", "RELEASENOTES_BARE.md"},
	}

	for _, testCase := range testCases {
		got := releasenotes.ReadFromFile(testCase.changelog)
		want := mustReadTestData(testCase.releaseNotes)
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
