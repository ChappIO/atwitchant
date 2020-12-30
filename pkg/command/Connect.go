package command

import (
	"atwitchant/pkg/twitch"
	"flag"
	"log"
	"os"
	"os/signal"
)

var profile = ""
var noCommon = false

var Connect = &Command{
	Name:        "connect",
	Description: "Connect to your stream and enable the bot",
	Flags: func() {
		flag.StringVar(&profile, "profile", "default.json", "the path to the profile configuration")
		flag.BoolVar(&noCommon, "no-common", false, "skip loading the common.json configuration")
	},
	Run: func() {
		api := twitch.LoadTwitch()
		if api.Token == "" {
			log.Println("Run the login command first")
			os.Exit(1)
			return
		}

		api.Chat.OnMessage(twitch.CommandPrivMsg, func(msg twitch.ChatMessage) {
			api.Chat.SendMessage("Hello " + msg.Tags["display-name"])
		})

		// wait for kill signal
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
	},
}
