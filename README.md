# Ocha

Journaling in the comfort of your terminal emulator.

---

ocha is supposed to be a cross-platform obsidian-like app in the terminal, where you an stores notes, lists, passwords & ssh-keys.

## Todo

-   buckets?

*   Note taking
*   Lists / Bullet points (toggle)
*   Password manager
*   ssh-keys

### CLI

-   [x] notes (maccha)
-   [ ] boards (sencha)
-   [ ] vault (bancha)
-   [ ] backup
-   [ ] config

### TUI

-   [x] Listing
-   [x] Editing
-   [x] Creating
-   [x] Renaming
-   [x] Markdown view

## Features

## Implementation

### CLI

-   Cobra
-   urfave/cli

### Storage

#### Options

-   Postgres + gorm
-   boltdb <- worth looking into
-   gorm + sqlite

### TUI

#### Options

-   bubbletea

## CLI

### Basic usage

```
ocha <command> [flags]
```

### Notes

```

ocha notes
ocha n

```

#### Subcommands

```

ocha notes list
ocha notes ls

```

List all of the notes (titles) currently saved in the database (optionally make them checkable to edit) it must return an id to the user

```

ocha notes create <name>
ocha notes c <name>

```

Create a new note titled `<name>`

```

ocha notes | n edit

```

it will display a list of all of the notes to the user to choose from

```

ocha notes | n delete --id <id>

```

delte a note using it's `<id>`m if hte id is not passed, the user will be prompted to select a note to delete

```

ocha notes search

```

filter out notes by title
