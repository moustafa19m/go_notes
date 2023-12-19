package journal

import (
	"fmt"
	"sort"
	"strings"

	j "github.com/moustafa19m/go_notes/pkg/journal"
)

// define a new interface for journal
type Client interface {
	Read(id int) (*j.Journal, error)
	Save(j *j.Journal) error
	Delete(id int) error
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
func (s *Service) Create(title string, content string, tags string) (*j.Journal, error) {
	tagsSlice := strings.Split(tags, ",")
	journal, err := j.NewJournal(title, content, tagsSlice)
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

// Delete a journal
func (s *Service) Delete(id int) error {
	return s.client.Delete(id)
}

// filter by title
func (s *Service) Filter(title string) ([]*j.Journal, error) {
	journals, err := s.client.ListAll()
	if err != nil {
		return nil, err
	}

	var filteredJournals []*j.Journal
	for _, journal := range journals {
		search := strings.ToLower(title)
		currenTitle := strings.ToLower(journal.Title)

		if strings.Contains(currenTitle, search) {
			filteredJournals = append(filteredJournals, journal)
		}
	}
	if len(filteredJournals) == 0 {
		return nil, fmt.Errorf("no journals found with title matches to (%s)", title)
	}
	return filteredJournals, nil
}

func (s *Service) SortAsc() ([]*j.Journal, error) {
	journals, err := s.client.ListAll()
	if err != nil {
		return nil, err
	}

	sort.Slice(journals, func(i, j int) bool {
		str1 := strings.Join([]string{journals[i].Title, journals[i].Content}, "")
		str2 := strings.Join([]string{journals[j].Title, journals[j].Content}, "")
		return strings.Compare(str1, str2) < 0
	})

	return journals, nil
}

func (s *Service) SortDesc() ([]*j.Journal, error) {
	journals, err := s.client.ListAll()
	if err != nil {
		return nil, err
	}

	sort.Slice(journals, func(i, j int) bool {
		str1 := strings.Join([]string{journals[i].Title, journals[i].Content}, "")
		str2 := strings.Join([]string{journals[j].Title, journals[j].Content}, "")
		return strings.Compare(str1, str2) > 0
	})

	return journals, nil
}

func (s *Service) AddTags(id int, tags string) (*j.Journal, error) {
	journal, err := s.client.Read(id)
	if err != nil {
		return nil, err
	}

	journal.AddTags(cleanup(tags))
	return journal, s.client.Save(journal)
}

func (s *Service) Analyze() (map[string]int, error) {
	journals, err := s.client.ListAll()
	if err != nil {
		return nil, nil
	}
	words := []string{}
	for _, journal := range journals {
		words = append(words, strings.Split(journal.Content, " ")...)
		words = append(words, strings.Split(journal.Title, " ")...)
	}

	wordsMap := make(map[string]int)
	for _, word := range words {
		wordsMap[word]++
	}

	return wordsMap, nil
}

func cleanup(tags string) []string {
	tagsSlice := strings.Split(tags, ",")
	var cleanedTags []string
	for _, tag := range tagsSlice {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}
		cleanedTags = append(cleanedTags, tag)
	}
	return cleanedTags
}
