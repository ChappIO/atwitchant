package command

import (
	"log"
	"os"
	"strings"
)

var All = []*Command{
	Login,
}

func GetCommand(name string) *Command {
	for _, command := range All {
		if strings.ToLower(command.Name) == strings.ToLower(name) {
			return command
		}
	}

	log.Printf("unknown command '%s' run 'twitchy help' to see all commands", name)
	os.Exit(1)
	return nil
}
