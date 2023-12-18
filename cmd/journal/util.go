package main

import (
	"encoding/json"
	"fmt"
	"os"

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
