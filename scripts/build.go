package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		args = append(args, "usage")
	}

	var err error

	switch args[0] {
	case "bin/axis":
		err = build()
		break
	case "clean":
		err = delete("bin")
		break
	default:
		err = usage()
		break
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func build() error {
	goos := os.Getenv("OS")
	goarch := os.Getenv("ARCH")
	tag := strings.ReplaceAll(os.Getenv("GITHUB_REF"), "refs/tags/", "")
	fmt.Printf("tag: %s\n", tag)
	out := fmt.Sprintf("./bin/axis_%s_%s", goos, goarch)
	ldflags := "-X 'github.com/FantasticFiasco/axis-cli/internal/build.Version=TODO' " +
		fmt.Sprintf("-X 'github.com/FantasticFiasco/axis-cli/internal/build.ReleaseUrl=%s' ", releaseURL(tag)) +
		fmt.Sprintf("-X 'github.com/FantasticFiasco/axis-cli/internal/build.Date=%s'", date())

	_, err := run(
		"go",
		"build",
		"-o",
		out,
		"-v",
		"-ldflags="+ldflags,
		"./cmd/axis")

	return err
}

func releaseURL(tag string) string {
	url := "https://github.com/FantasticFiasco/axis-cli/releases"
	if tag != "" {
		url = fmt.Sprintf("%s/tag/%s", url, tag)
	}
	return url
}

func date() string {
	t := time.Now()
	return t.Format("2006-01-02")
}

func run(exe string, args ...string) ([]byte, error) {
	fmt.Printf("%s %s\n", exe, strings.Join(args, " "))
	out, err := exec.Command(exe, args...).CombinedOutput()
	fmt.Println(string(out))
	return out, err
}

func delete(targets ...string) error {
	fmt.Printf("Delete %s\n", strings.Join(targets, ", "))
	for _, target := range targets {
		if err := os.RemoveAll(target); err != nil {
			return err
		}
	}
	return nil
}

func usage() error {
	return errors.New(
		strings.Join([]string{
			"Usage: go run scripts/build.go <task>",
			"",
			"Known tasks are:",
			"",
			"  bin/axis:",
			"    Builds the main executable.",
			"    Supported environment variables:",
			"    - OS",
			"    - ARCH",
			"    - GITHUB_REF",
			"",
			"  clean:",
			"    Deletes all built files.",
			"",
		},
			"\n"),
	)
}
