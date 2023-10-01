package sn

const RefTypeTagToParentTag = "TagToParentTag"

type Reference struct {
	UUID          string `json:"uuid"`
	ContentType   string `json:"content_type"`
	ReferenceType string `json:"reference_type,omitempty"`
}
