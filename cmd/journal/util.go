package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	j "github.com/moustafa19m/apple_interview/pkg/journal"

	service "github.com/moustafa19m/apple_interview/service/journal"
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

func runCreate(js *service.Service, title string, content string) {
	// create a new journal
	_, err := js.Create(title, content)
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
