package main

import (
	"atwitchant/pkg/command"
	"flag"
	"log"
	"os"
)

var printHelp = false

func main() {
	// find the commands to be able to select the right implementation
	var selectedCommand *command.Command
	if len(os.Args) >= 2 {
		selectedCommand = command.GetCommand(os.Args[1])
		os.Args = os.Args[1:]
	}
	if selectedCommand == nil {
		selectedCommand = command.Help
	}

	flag.BoolVar(&printHelp, "help", false, "print this message")
	if selectedCommand.Flags != nil {
		selectedCommand.Flags()
	}
	flag.Parse()

	if printHelp {
		log.Print("\n" + selectedCommand.Description + "\n\nflags:")
		flag.PrintDefaults()
		return
	}

	if selectedCommand.Validate != nil {
		selectedCommand.Validate()
		flag.CommandLine.Usage()
	}
	selectedCommand.Run()
}
