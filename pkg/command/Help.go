package command

import (
	"flag"
	"log"
)

var Help = &Command{
	Name: "help",
	Description: "Here's how to use your favorite twitch assistant!",
	Flags: func() {

	},
	Run: func() {
		log.Println("welcome to twitchy")
		log.Println("here's a list of commands:")

		for _, command := range All {
			log.Printf("\t- %s: %s", command.Name, command.Description)
		}

		log.Println("flags:")
		flag.PrintDefaults()
	},
}
