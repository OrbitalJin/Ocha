package app

import (
	notes "github.com/orbitaljin/ocha/cmd"
	"github.com/orbitaljin/ocha/internal/store"
	"github.com/urfave/cli/v2"
)


var commands []*cli.Command = make([]*cli.Command, 0)

func Cog(db *store.DB) []*cli.Command {
	commands = append(commands, notes.Handler(db))
	return commands
}

func New(db *store.DB) *cli.App {
	return &cli.App{
		Name: "Gault",
		Usage: "Cli app like osi lmaoo",
		Commands: Cog(db),
	}
}