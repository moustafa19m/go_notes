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
	}, nil
}

func (jc Client) Read(title string) (*j.Journal, error) {
	return nil, nil
}

func (jc Client) Save(j *j.Journal) error {
	file, err := os.OpenFile(jc.filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	jc.journals = append(jc.journals, j)

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
