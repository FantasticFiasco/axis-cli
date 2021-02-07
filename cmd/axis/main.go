package main

// The following variables are set by goreleaser during CD
var version = "<version>"
var commit = "<git sha>"
var date = "<date>"

func main() {
	// TODO: Create basic configuration of the CLI parser
	// TODO: Re-implement --version (https://github.com/urfave/cli/blob/master/docs/v2/manual.md#version-flag)

	//versionFlag := flag.Bool("version", false, "Show axis version")
	//searchFlag := flag.Bool("search", false, "Search for devices on the network")
	//flag.Parse()
	//
	//var err error
	//
	//if *versionFlag == true {
	//	commands.Version(version, commit, date)
	//} else if *searchFlag {
	//	err = commands.Search()
	//} else {
	//	flag.Usage()
	//}

	//if err != nil {
	//	fmt.Fprint(os.Stderr, err)
	//	os.Exit(1)
	//}
}
