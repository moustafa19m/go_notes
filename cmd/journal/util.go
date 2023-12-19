package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"

	j "github.com/moustafa19m/go_notes/pkg/journal"

	service "github.com/moustafa19m/go_notes/service/journal"
)

func printHelp() {
	// Print the help message
	// supported commands are "create", "list"
	// create takes a title as an argument
	// list takes no arguments
	helpMessage := `
	Usage: journal [command] [arguments]
	The commands are:
	  * create:	creates a new journal
	     ./journal --create "Content of the journal here" --title "Title here"
	  * list:  	lists all journals
	     ./journal --list
	`
	println(helpMessage)
}

func runCreate(js *service.Service, title string, content string, tags string) {
	// create a new journal
	_, err := js.Create(title, content, tags)
	if err != nil {
		fmt.Printf("failed to create new journal: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Journal with title (%s) created successfully\n", title)
}

func runList(js *service.Service) {
	// list all journals
	journals, err := js.ListAll()
	if err != nil {
		fmt.Printf("failed to list journals: %s\n", err.Error())
		os.Exit(1)
	}
	data, err := json.MarshalIndent(journals, "", "    ")
	if err != nil {
		fmt.Printf("failed to marshal journals: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(string(data))
}

func runTags(js *service.Service, tags string, id string) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("failed to convert id (%s) to int: %s\n", id, err.Error())
		os.Exit(1)
	}
	journal, err := js.AddTags(intId, tags)
	if err != nil {
		fmt.Printf("failed to add tags to journal: %s\n", err.Error())
		os.Exit(1)
	}
	data, err := json.MarshalIndent(journal, "", "    ")
	if err != nil {
		fmt.Printf("failed to marshal journals: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(string(data))
}

func runDelete(js *service.Service, id string) {
	// delete a journal
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("failed to convert id to int: %s\n", err.Error())
		os.Exit(1)
	}
	err = js.Delete(intId)
	if err != nil {
		fmt.Printf("failed to delete journal: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Journal with id (%d) deleted successfully\n", intId)
}

func runFilter(js *service.Service, title string) {
	// filter journals by title
	journals, err := js.Filter(title)
	if err != nil {
		fmt.Printf("failed to filter journals: %s\n", err.Error())
		os.Exit(1)
	}
	data, err := json.MarshalIndent(journals, "", "    ")
	if err != nil {
		fmt.Printf("failed to marshal journals: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(string(data))
}

func runSort(js *service.Service, sortOrder string) {
	var journals []*j.Journal
	var err error
	switch sortOrder {
	case "desc":
		journals, err = js.SortDesc()
	case "asc", "":
		journals, err = js.SortAsc()
	default:
		fmt.Printf("invalid sort order: %s\n", sortOrder)
	}
	if err != nil {
		fmt.Printf("failed to sort journals: %s\n", err.Error())
		os.Exit(1)
	}
	data, err := json.MarshalIndent(journals, "", "    ")
	if err != nil {
		fmt.Printf("failed to marshal journals: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(string(data))
}

func runAnalyze(js *service.Service) {
	// analyze a journal
	analysis, err := js.Analyze()
	if err != nil {
		fmt.Printf("failed to analyze journals: %s\n", err.Error())
		os.Exit(1)
	}

	type out struct {
		Count int    `json:"count"`
		Word  string `json:"word"`
	}

	outSlice := []out{}
	for k, v := range analysis {
		outSlice = append(outSlice, out{v, k})
	}
	sort.Slice(outSlice, func(i, j int) bool {
		return outSlice[i].Count > outSlice[j].Count
	})
	data, err := json.MarshalIndent(outSlice, "", "    ")
	if err != nil {
		fmt.Printf("failed to marshal journals: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(string(data))
}
