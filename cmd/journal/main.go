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
	{
		Name:   "filter",
		Help:   "filter journals by title",
		IsBool: false,
	},
	{
		Name:   "sort",
		Help:   "sort journals by title, takes value asc or desc",
		IsBool: false,
	},
	{
		Name:   "tags",
		Help:   "tags to add to the journal",
		IsBool: false,
	},
	{
		Name:   "id",
		Help:   "id of the journal to process",
		IsBool: false,
	},
}

// .journal --tags "a,b,c" --id=1

// .journal --sort asc
// .journal --filter "Title of the journal to filter"

// ./journal --delete "Title of the journal to delete"

//.journal --filter="title"
//.journal --filter="a, b, c"
//.journal --fliter --title="title" --tags="a,b,c"

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
		runCreate(js, args["title"], args["create"], args["tags"])
	case "list":
		// list all journals
		runList(js)
	case "help":
		// print help
		printHelp()
	case "delete":
		// delete a journal
		runDelete(js, args["delete"])
	case "filter":
		// filter journals by title
		runFilter(js, args["filter"])
	case "sort":
		// sort journals by title
		runSort(js, args["sort"])
	case "tags":
		// add tags to a journal
		runTags(js, args["tags"], args["id"])
	default:
		fmt.Println("Invalid command, run --help to see the list of commands")
	}
	// exit program
}
