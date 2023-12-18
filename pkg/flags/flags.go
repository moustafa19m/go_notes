package arg_parser

import (
	"flag"
	"os"
	"strings"
)

type FlagDef struct {
	Name   string
	Help   string
	IsBool bool
}

// ParseArgs parses the command line arguments and returns the command and the arguments
func ParseFlags(flagDefs []FlagDef) (string, map[string]string, error) {
	// Parse the command line arguments
	// return a map of key value pairs
	for _, f := range flagDefs {
		if !f.IsBool {
			flag.String(f.Name, "", f.Help)
		} else {
			flag.Bool(f.Name, false, f.Help)
		}
	}

	flag.Parse()

	argsMap := make(map[string]string)
	// visit all flags
	flag.VisitAll(func(f *flag.Flag) {
		// get flag key and value
		// fmt.Printf("flag: %s, value: %s\n", f.Name, f.Value)
		argsMap[clean(f.Name)] = f.Value.String()
	})

	return clean(os.Args[1]), argsMap, nil
}

func clean(name string) string {
	for {
		if strings.HasPrefix(name, "-") {
			name = name[1:]
		} else {
			break
		}
	}
	name = strings.Split(name, "=")[0]
	return name
}
