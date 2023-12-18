package journal_os

import (
	"encoding/json"
	"fmt"
	"os"

	j "github.com/moustafa19m/apple_interview/pkg/journal"
)

type Client struct {
	filename string
	journals []*j.Journal
	latestId int
}

func NewJournalOs(filename string) (Client, error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println("file does not exist, will create a new one when saving")
		return Client{
			filename: filename,
		}, nil
	}
	journals, err := readSavedJournals(filename)
	if err != nil {
		return Client{}, fmt.Errorf("failed to read saved journals: %s", err.Error())
	}

	return Client{
		filename: filename,
		journals: journals,
		latestId: journals[len(journals)-1].Id,
	}, nil
}

func (jc Client) Delete(id int) error {
	for i, journal := range jc.journals {
		if journal.Id == id {
			jc.journals = append(jc.journals[:i], jc.journals[i+1:]...)
			return jc.saveFile()
		}
	}
	return fmt.Errorf("journal with id (%d) does not exist", id)
}

func (jc Client) JournalExists(id int) bool {
	for _, journal := range jc.journals {
		if journal.Id == id {
			return true
		}
	}
	return false
}

func (jc Client) Read(id int) (*j.Journal, error) {
	for _, journal := range jc.journals {
		if journal.Id == id {
			return journal, nil
		}
	}
	return nil, fmt.Errorf("journal with id (%d) does not exist", id)
}

func (jc Client) createId() int {
	jc.latestId += 1
	return jc.latestId
}

func (jc Client) Save(j *j.Journal) error {
	if j.Id == 0 {
		j.Id = jc.createId()
		jc.journals = append(jc.journals, j)
	} else {
		for i, journal := range jc.journals {
			if journal.Id == j.Id {
				jc.journals[i] = j
			}
		}
	}
	return jc.saveFile()
}

func (jc Client) saveFile() error {
	file, err := os.OpenFile(jc.filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(jc.journals)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (jc Client) ListAll() ([]*j.Journal, error) {
	return jc.journals, nil
}

// file is a json file
// in the following format:
//
// [
//
//		{
//			"title": "title1",
//			"content": "content1"
//		},
//		{
//			"title": "title2",
//			"content": "content2"
//	  }
//
// ]
func readSavedJournals(filename string) ([]*j.Journal, error) {
	// Read the file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Decode the JSON data
	var journals []*j.Journal
	err = json.Unmarshal(data, &journals)
	if err != nil {
		return nil, err
	}

	return journals, nil
}
