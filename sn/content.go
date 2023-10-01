package sn

const (
	ContentTypeNote = "Note"
	ContentTypeTag  = "Tag"
)

type Content struct {
	Title      string      `json:"title"`
	Text       string      `json:"text"`
	References []Reference `json:"references"`
}
