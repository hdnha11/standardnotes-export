package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hdnha11/standardnotes-export/sn"
)

func main() {
	var (
		backupFilePath string
		outputFolder   string
	)

	flag.StringVar(&backupFilePath, "f", "", "Standard Notes Backup and Import File")
	flag.StringVar(&outputFolder, "o", "./md", "Markdown Output Folder")
	flag.Parse()

	snBackup := sn.OpenBackup(backupFilePath)

	fmt.Println("Migrating")
	fmt.Printf("- %d notes\n", len(snBackup.Notes()))
	fmt.Printf("- %d tags\n", len(snBackup.Tags()))

	mds := snBackup.ExportMarkdown()
	writeMarkdownFiles(outputFolder, mds)

	fmt.Printf("- exported %d Markdown files to %s\n", len(mds), outputFolder)
	fmt.Println("- done!")
}

func writeMarkdownFiles(outputPath string, mds []sn.MarkdownExport) {
	ensureDir(outputPath)
	for _, md := range mds {
		if err := os.WriteFile(
			filepath.Join(outputPath, escapeSeparator(md.FileName)),
			[]byte(md.Content),
			0660,
		); err != nil {
			log.Fatal(err)
		}
	}
}

func escapeSeparator(filename string) string {
	return strings.ReplaceAll(filename,
		fmt.Sprintf("%c", filepath.Separator), "-")
}

func ensureDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0750); err != nil {
			log.Fatal(err)
		}
	}
}
