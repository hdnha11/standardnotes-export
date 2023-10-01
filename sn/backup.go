package sn

import (
	"encoding/json"
	"log"
	"os"
	"sort"
)

type Backup struct {
	Items []Item `json:"items"`
}

type Item struct {
	ContentType string  `json:"content_type"`
	Content     Content `json:"content"`
	Deleted     bool    `json:"deleted"`
	UUID        string  `json:"uuid"`
}

func OpenBackup(path string) Backup {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var bak Backup
	if err := json.Unmarshal(content, &bak); err != nil {
		log.Fatal(err)
	}

	return bak
}

func (b Backup) Notes() []Note {
	tags := b.Tags()
	lookup := createReverseNotesLookupFromTags(tags)

	var notes []Note
	for _, item := range b.Items {
		if item.ContentType == ContentTypeNote {
			notes = append(notes, Note{
				UUID:  item.UUID,
				Title: item.Content.Title,
				Text:  item.Content.Text,
				Tags:  getTags(item.UUID, lookup),
			})
		}
	}
	return notes
}

func (b Backup) Tags() []Tag {
	var (
		tags       []Tag
		tagsLookup = make(map[string]*Tag)
	)

	for _, item := range b.Items {
		if item.ContentType == ContentTypeTag {
			tag := Tag{
				UUID:       item.UUID,
				Title:      item.Content.Title,
				NoteUUIDs:  getNoteUUIDs(item.Content.References),
				parentUUID: getParentTagUUID(item.Content.References),
			}
			tags = append(tags, tag)
			tagsLookup[tag.UUID] = &tag
		}
	}

	// Assign parent tag
	for i := range tags {
		tag := &tags[i]
		if tag.parentUUID != nil {
			tag.Parent = tagsLookup[*tag.parentUUID]
		}
	}

	return tags
}

func getParentTagUUID(refs []Reference) *string {
	for _, ref := range refs {
		if ref.ContentType == ContentTypeTag && ref.ReferenceType == RefTypeTagToParentTag {
			parentUUID := ref.UUID
			return &parentUUID
		}
	}
	return nil
}

func getNoteUUIDs(refs []Reference) []string {
	var uuids []string
	for _, ref := range refs {
		if ref.ContentType == ContentTypeNote {
			uuids = append(uuids, ref.UUID)
		}
	}
	return uuids
}

func createReverseNotesLookupFromTags(tags []Tag) map[string][]Tag {
	lookup := make(map[string][]Tag)
	for _, tag := range tags {
		for _, noteUUID := range tag.NoteUUIDs {
			lookup[noteUUID] = append(lookup[noteUUID], tag)
		}
	}
	return lookup
}

func getTags(noteUUID string, lookup map[string][]Tag) []Tag {
	tags := lookup[noteUUID]
	sort.Slice(tags, func(i, j int) bool {
		return tags[i].FullTitle() < tags[j].FullTitle()
	})
	return removeDup(tags)
}

func removeDup[T any](xs []T) []T {
	// TODO
	return xs
}
