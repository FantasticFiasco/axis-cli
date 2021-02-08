package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var searchCommand = &cli.Command{
	Name: "search",
	Action: func(c *cli.Context) error {
		fmt.Println("TODO: Implement")
		return nil
	},
}
