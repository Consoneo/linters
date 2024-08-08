package main

import (
	"fmt"
	"os"

	"github.com/Consoneo/linters/src/engine"
	"github.com/charmbracelet/lipgloss"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {

	analyse := engine.Analyse{}

	app := &cli.App{
		Name:  "linters",
		Usage: "Lint and share your linting rules across projects",
		Authors: []*cli.Author{
			{
				Name:  "Consoneo",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "lint",
				Aliases: []string{"l"},
				Usage:   "lint your projects",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:     "verbose",
						Aliases:  []string{"v"},
						Usage:    "Enable verbose mode",
						Category: "Global options",
					},
				},
				Action: func(cCtx *cli.Context) error {
					
					if cCtx.Bool("verbose") {
						log.SetLevel(log.DebugLevel)
					}

					err := analyse.Lint()
					if err == nil {
						// at the bottom of the screen
						style := lipgloss.NewStyle().Background(lipgloss.Color("#00FF00")).Foreground(lipgloss.Color("#000000")).AlignVertical(lipgloss.Bottom).MarginTop(2)
						fmt.Fprintf(os.Stdout, "%s\n", style.Render("Success"))
					}

					analyse.ListReports()
					return err
				},
			},
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "Initialize config file",
				Action: func(cCtx *cli.Context) error {
					analyse.InitConfig()
					return nil
				},
			},
			{
				Name:    "rules",
				Aliases: []string{"r"},
				Usage:   "List all available rules",
				Action: func(cCtx *cli.Context) error {
					analyse.ListRules()
					return nil
				},
			},
			{
				Name:    "install",
				Aliases: []string{"n"},
				Usage:   "Initialize git hooks in current folder",
				Action: func(cCtx *cli.Context) error {
					analyse.Install()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		style := lipgloss.NewStyle().Background(lipgloss.Color("#FF0000")).Foreground(lipgloss.Color("#000000"))
		fmt.Fprintf(os.Stderr, "%s\n", style.Render(err.Error()))
	}

}
