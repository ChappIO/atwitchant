package command

import (
	"atwitchant/pkg/config"
	"encoding/json"
	"flag"
	"log"
	"os"
)

var New = &Command{
	Name:        "new",
	Description: "Initialize a new profile",
	Flags: func() {
		flag.StringVar(&profile, "profile", "default.json", "the path to the profile configuration")
	},
	Run: func() {
		profileData := config.Profile{Triggers: map[string]config.Trigger{
			"ping_command": {
				Comment: "This is just an example trigger which gets fired when a viewer sends '!ping'",
				Match:   config.Match{
					Comment: "When the message is '!ping'",
					Message: config.MustCompile("^!ping$"),
				},
				Action: "",
			},
		}}
		file, err := os.Create(profile)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		enc := json.NewEncoder(file)
		enc.SetIndent("", "  ")
		enc.Encode(&profileData)
		log.Printf("saved new profile to '%s'", profile)
	},
}