# Standard Notes Markdown Export

A simple tool to help migrate from Standard Notes backup files to other Markdown-based note-taking apps like Obsidian.

This tool takes a straightforward approach - keeping all note files flattened, with the file name being the note's title followed by `.md`.

It also retains the original tags from each note by adding a front matter section to the top of the file if tags exist.

## Build

```bash
$ go build main.go
```

## Usage

```bash
$ ./main -h

Usage of ./main:
  -f string
        Standard Notes Backup and Import File
  -o string
        Markdown Output Folder (default "./md")
```

Examples:

```bash
$ ./main -f ~/Downloads/backup.txt
```
