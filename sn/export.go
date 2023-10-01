package sn

import (
	"fmt"
	"strings"
)

type MarkdownExport struct {
	FileName string
	Content  string
}

func (b Backup) ExportMarkdown() []MarkdownExport {
	var mds []MarkdownExport

	notes := b.Notes()
	for _, note := range notes {
		mds = append(mds, MarkdownExport{
			FileName: fmt.Sprintf("%s.md", note.Title),
			Content:  markdownContent(note),
		})
	}

	return mds
}

func markdownContent(note Note) string {
	var frontMatter string
	if len(note.Tags) > 0 {
		tags := strings.Join(note.TagTitles(), ", ")
		frontMatter = fmt.Sprintf("---\ntags: %s\n---", tags)
	}
	if frontMatter == "" {
		return note.Text
	}

	return fmt.Sprintf("%s\n\n%s", frontMatter, note.Text)
}
