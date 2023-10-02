package sn

import (
	"strings"

	"github.com/hdnha11/standardnotes-export/slices"
)

type Tag struct {
	UUID      string
	Title     string
	Parent    *Tag
	NoteUUIDs []string

	parentUUID *string
}

func (t Tag) FullTitle() string {
	var titles []string

	cur := &t
	for cur != nil {
		titles = append(titles, cur.Title)
		cur = cur.Parent
	}

	slices.Reverse(titles)
	return strings.Join(titles, "/")
}
