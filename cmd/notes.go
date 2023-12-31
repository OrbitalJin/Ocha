package notes

import (
	"fmt"

	"github.com/orbitaljin/ocha/internal/store"
	"github.com/orbitaljin/ocha/internal/store/schema"
	"github.com/urfave/cli/v2"
)

func Handler(db *store.DB) *cli.Command {
	return &cli.Command{
			Name:    "notes",
			Aliases: []string{"n"},
			Usage:   "manage your notes",
			Subcommands: subcommands(db),
		}
}

func subcommands(db *store.DB) []*cli.Command {
		subcommands := make([]*cli.Command, 0)
		subcommands = append(subcommands, create(db))
		subcommands = append(subcommands, list(db))
		return subcommands
}

// Idea: use create/edited at instead of description
func create(db *store.DB) *cli.Command {
	return &cli.Command{
		Name: "create",
		Aliases: []string{"c"},
		Usage: "create new note",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "title",
				Usage: "specify the title",
			},
		},
		Action: func(ctx *cli.Context) error {
			title := ctx.String("title")
			if title == "" {
				return fmt.Errorf("no title provided")
			}
			db.DB().Create(&schema.Note{
				ItemTitle: title,
			})
			fmt.Println("Created note:", title)
			return nil
		},
	}
}

func list(db *store.DB) *cli.Command {
	return &cli.Command {
		Name: "list",
		Aliases: []string{"ls"},
		Usage: "list all your notes",
		Flags: []cli.Flag {
			&cli.StringFlag {
				Name: "filter",
				Usage: "filter out the notes",
			},
		},
		Action: func(ctx *cli.Context) error {
			var records []schema.Note
			db.DB().Find(&records)
			for record := range records {
				fmt.Println(record)
			}
			return nil
		},
	}
}