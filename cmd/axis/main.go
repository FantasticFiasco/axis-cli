package main

import (
	"github.com/FantasticFiasco/axis-cli/internal/console"
	"github.com/urfave/cli/v2"
	"os"
)

// The following variables are set by goreleaser during CD
var version = "<version>"
var commit = "<git sha>"
var date = "<date>"

func main() {
	app := cli.App{
		Name:    "axis",
		Usage:   "axis is a tool for managing devices from Axis Communications.",
		Version: version,
		Commands: []*cli.Command{
			searchCommand,
		},
	}

	cli.VersionPrinter = versionPrinter

	err := app.Run(os.Args)
	if err != nil {
		console.Fatal(err)
	}

	//searchFlag := flag.Bool("search", false, "Search for devices on the network")
}

func versionPrinter(c *cli.Context) {
	url := "https://github.com/FantasticFiasco/axis-cli/releases"
	if version != "<version>" {
		url += "/tag/" + version
	}

	console.Printf("axis %s\n", version)
	console.Printf("commit:  %s\n", commit)
	console.Printf("release: %s\n", url)
	console.Printf("date:    %s\n", date)
}
