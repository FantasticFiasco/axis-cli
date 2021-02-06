package commands

import (
	"fmt"
)

func Version(version, commit, date string) {
	fmt.Printf("axis %s\n", version)
	fmt.Printf("commit:  %s\n", commit)
	fmt.Printf("release: %s\n", releaseURL(version))
	fmt.Printf("date:    %s\n", date)
}

func releaseURL(version string) string {
	url := "https://github.com/FantasticFiasco/axis-cli/releases"
	if version != "<version>" {
		url += "/tag/" + version
	}
	return url
}
