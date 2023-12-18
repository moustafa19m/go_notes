package journal

import "errors"

type Journal struct {
	Id      int      `json:"id"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
	Title   string   `json:"title"`
}

var ErrEmptyTitle = errors.New("empty title")
var ErrEmptyContent = errors.New("empty content")

// Create a new jouranl with content and title
func NewJournal(title string, content string, tags []string) (*Journal, error) {
	if title == "" {
		return nil, ErrEmptyTitle
	}
	if content == "" {
		return nil, ErrEmptyContent
	}
	return &Journal{Title: title, Content: content, Tags: tags}, nil
}

// Add the content to the journal
func (j *Journal) AddContent(content string) {
	// append to content
	j.Content += content
}

// Add tasgs to the journal
func (j *Journal) AddTags(tags []string) {
	j.Tags = append(j.Tags, tags...)
}
