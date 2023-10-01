package sn

type Note struct {
	UUID  string
	Title string
	Text  string
	Tags  []Tag
}

func (n Note) TagTitles() []string {
	var titles []string
	for _, tag := range n.Tags {
		titles = append(titles, tag.FullTitle())
	}
	return titles
}
