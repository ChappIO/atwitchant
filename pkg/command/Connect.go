package command

import (
	"atwitchant/pkg/config"
	"atwitchant/pkg/twitch"
	"bytes"
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"os"
	"os/signal"
	"sort"
)

var profile = ""
var noCommon = false

func loadProfile(profile *config.Profile, fileName string) {
	log.Printf("loading %s", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(profile)
	if err != nil {
		panic(err)
	}
}

type messageMatch struct {
	Score  int
	Action string
}

var Connect = &Command{
	Name:        "connect",
	Description: "Connect to your stream and enable the bot",
	Flags: func() {
		flag.StringVar(&profile, "profile", "default.json", "the path to the profile configuration")
		flag.BoolVar(&noCommon, "no-common", false, "skip loading the common.json configuration")
	},
	Run: func() {
		profileData := config.Profile{}

		loadProfile(&profileData, "common.json")
		loadProfile(&profileData, profile)

		log.Printf("%+v", profileData)

		api := twitch.LoadTwitch()
		if api.Token == "" {
			log.Println("Run the login command first")
			os.Exit(1)
			return
		}

		api.Chat.OnMessage(twitch.CommandPrivMsg, func(msg twitch.ChatMessage) {
			matches := []messageMatch{}

			for _, trigger := range profileData.Triggers {
				score := trigger.Check(&msg)
				if score >= 0 {
					matches = append(matches, messageMatch{
						Score:  score,
						Action: trigger.Action,
					})
				}
			}

			sort.Slice(matches, func(i, j int) bool {
				return matches[i].Score < matches[j].Score
			})

			sender := msg.Tags["display-name"]
			if len(matches) > 0 {
				log.Printf("%s said: %s", sender, msg.Body)
			}
			for _, match := range matches {
				log.Printf("%s trigged action %s", sender, match.Action)
				if action, ok := profileData.Actions[match.Action]; !ok {
					log.Printf("%s does not exist", match.Action)
				} else {
					runAction(&api, &action, &msg)
				}
			}
		})

		// wait for kill signal
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
	},
}

func runAction(api *twitch.Integration, action *config.Action, msg *twitch.ChatMessage) {
	if action.SendMessage != nil {
		runSendMessage(api, action.SendMessage, msg)
	}
}

func runSendMessage(api *twitch.Integration, action *config.SendMessageAction, msg *twitch.ChatMessage) {
	var tpl bytes.Buffer
	tmpl, err := template.New("").Parse(action.Template)
	if err != nil {
		log.Printf("error: failed to compile template: %s", err)
		return
	}
	err = tmpl.Execute(&tpl, map[string]interface{}{
		"msg": msg,
	})
	if err != nil {
		log.Printf("error: failed to compile message: %s", err)
		return
	}

	api.Chat.SendMessage(tpl.String())
}
