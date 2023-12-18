package journal

import (
	j "github.com/moustafa19m/apple_interview/pkg/journal"
)

// define a new interface for journal
type Client interface {
	Read(title string) (*j.Journal, error)
	Save(j *j.Journal) error
	ListAll() ([]*j.Journal, error)
}

// define a new struct for journal
type Service struct {
	// journals []j.Journal
	client Client
}

// Create a new journal service
func NewJournalService(jc Client) *Service {
	return &Service{
		client: jc,
	}
}

// Create a new journal
func (s *Service) Create(title string, content string) (*j.Journal, error) {
	journal, err := j.NewJournal(title, content)
	if err != nil {
		return nil, err
	}
	err = s.client.Save(journal)
	if err != nil {
		return nil, err
	}
	return journal, nil
}

// List all journals
func (s *Service) ListAll() ([]*j.Journal, error) {
	return s.client.ListAll()
}
