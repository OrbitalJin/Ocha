package cmd

import (
	"fmt"

	"github.com/orbitaljin/ocha/internal"
	"github.com/orbitaljin/ocha/internal/notepad"
	"github.com/orbitaljin/ocha/internal/store"
	"github.com/orbitaljin/ocha/internal/store/schema"
	"github.com/orbitaljin/ocha/utils"
	"github.com/urfave/cli/v2"
)

func NotesHandler(db *store.DB) *cli.Command {
	return &cli.Command{
			Name:    "notes",
			Aliases: []string{"maccha", "n"},
			Usage:   "manage your notes",
			Subcommands: subcommands(db),
		}
}

func subcommands(db *store.DB) []*cli.Command {
		subcommands := make([]*cli.Command, 0)
		subcommands = append(subcommands, create(db))
		subcommands = append(subcommands, list(db))
		subcommands = append(subcommands, import_(db))
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
		Action: func(ctx *cli.Context) error {
			var records []schema.Note
			db.DB().Find(&records)
			app := notepad.New(db, records)
			internal.Launch(app)
			return nil
		},
	}
}

func import_(db *store.DB) *cli.Command {
	return &cli.Command{
		Name: "import",
		Aliases: []string{"i"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "path",
				Usage: "specify the file path",
			},
		},
		Action: func (ctx *cli.Context) error {
			path := ctx.String("path")
			if path == "" {
				return fmt.Errorf("no file provided")
			}
			name, content, err := utils.Read(path)
			if err != nil {
				return err
			}
			db.DB().Create(&schema.Note{
				ItemTitle: name,
				Content: content,
			})
			fmt.Println("Imported note:", name)
			return nil
		},		
	}
}