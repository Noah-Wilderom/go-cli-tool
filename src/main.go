package main

import (
	"github.com/Noah-Wilderom/go-cli/src/commands"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Name:                 "CLI",
		Version:              "v0.0.1",
		EnableBashCompletion: true,
		Suggest:              true,
		Commands: []*cli.Command{
			{
				Name:    "production",
				Aliases: []string{"p"},
				Usage:   "production resources",
				Action: func(cCtx *cli.Context) error {
					commands.RunProduction(cCtx)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
