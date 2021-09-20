package cli

import (
	"flag"
	"fmt"
	"os"
)

var (
	allCommand []command
)

type command struct {
	Name        string
	Description string

	Run func([]string) error
}

func Usage() {
	_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()

	_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Sub commands:\n")

	for i := range allCommand {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "  %s: %s\n", allCommand[i].Name, allCommand[i].Description)
	}

}

func Dispatch(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("atleast one arg is required")
	}

	sub := args[0]
	for i := range allCommand {
		if sub == allCommand[i].Name {
			return allCommand[i].Run(args)
		}
	}

	flag.Usage()
	return fmt.Errorf("invalid command")
}

func addCommand(name, description string, run func([]string) error) {
	allCommand = append(allCommand, command{
		Name:        name,
		Description: description,
		Run:         run,
	})
}
