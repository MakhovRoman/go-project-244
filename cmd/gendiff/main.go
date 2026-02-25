package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	(&cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output format",
			},
		},
	}).Run(context.Background(), os.Args)
}
