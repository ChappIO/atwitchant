package command

import (
	"atwitchant/pkg/twitch"
	"bufio"
	"log"
	"os"
)

var Chat = &Command{
	Name:        "chat",
	Description: "Start a chat from terminal (for testing purposes)",
	Run:         func() {
		api := twitch.LoadTwitch()
		if api.Token == "" {
			log.Println("Run the login command first")
			os.Exit(1)
			return
		}

		api.Chat.OnMessage(twitch.CommandPrivMsg, func(msg twitch.ChatMessage) {
			log.Printf("%s: %s", msg.Tags["display-name"], msg.Body)
		})

		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			line := scan.Text()
			api.Chat.SendMessage(line)
		}
	},
}
