package main

import (
	"fmt"
	"os"

	client "github.com/moustafa19m/apple_interview/client/journal_os"
	f "github.com/moustafa19m/apple_interview/pkg/flags"
	service "github.com/moustafa19m/apple_interview/service/journal"
)

const (
	OUTPUT_FILE = "/Users/moustafamakhlouf/Documents/apple_interview/saved/journals.json"
)

var flags = []f.FlagDef{
	{
		Name:   "create",
		Help:   "create a new journal, requires a title and content",
		IsBool: false,
	},
	{
		Name:   "title",
		Help:   "title of the journal",
		IsBool: false,
	},
	{
		Name:   "list",
		Help:   "list all journals",
		IsBool: true,
	},
	{
		Name:   "help",
		Help:   "print help",
		IsBool: true,
	},
	{
		Name:   "delete",
		Help:   "delete a journal, requires a title",
		IsBool: false,
	},
	// ./journal --delete "Title of the journal to delete"
}

func main() {

	// read args from command line using flags
	cmd, args, err := f.ParseFlags(flags)
	if err != nil {
		fmt.Println(err)
		printHelp()
		os.Exit(1)
	}
	// fmt.Printf("cmd: %s, args: %v\n", cmd, args)

	// create a new jouranl client using journal_os
	// create a new journal using journal serivce
	jc, err := client.NewJournalOs(OUTPUT_FILE)
	if err != nil {
		fmt.Printf("os client error: %s\n", err.Error())
		os.Exit(1)
	}
	js := service.NewJournalService(jc)

	// fmt.Println(cmd)
	// switch on the command
	switch cmd {
	case "create":
		// create a new journal
		runCreate(js, args["title"], args["create"])
	case "list":
		// list all journals
		runList(js)
	case "help":
		// print help
		printHelp()
	case "delete":
		// delete a journal
		runDelete(js, args["delete"])
	default:
		fmt.Println("Invalid command, run --help to see the list of commands")
	}
	// exit program
}
