package sn

import (
	"strings"
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

	reverse(titles)
	return strings.Join(titles, "/")
}

func reverse[T any](xs []T) {
	for i := 0; i < len(xs)/2; i++ {
		xs[i], xs[len(xs)-1-i] = xs[len(xs)-1-i], xs[i]
	}
}
