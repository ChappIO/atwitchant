package command

import (
	"atwitchant/pkg/open"
	"atwitchant/pkg/twitch"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

const loginPage = `
<!DOCTYPE html>
<html lang="en">
<body>
<p>This window should close itself<p>
<script>
const hash = window.location.hash.substring(1);
const params = new URLSearchParams(hash);
const token = params.get('access_token');
if(token) {
	fetch('/token', {
	  method: 'POST',
	  headers: {
		'Content-Type': 'text/plain',
	  },
	  body: token
	}).then(() => window.close());
}
</script>
</body>
</html>
`

var Login = &Command{
	Name:        "login",
	Description: "Authenticate with your twitch account",
	Flags: func() {

	},
	Run: func() {
		log.Println("authenticating with twitch...")
		waitForToken := make(chan string)

		mux := http.NewServeMux()
		mux.HandleFunc("/twitch_login.html", func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-Type", "text/html")
			_, _ = writer.Write([]byte(loginPage))
		})
		mux.HandleFunc("/token", func(writer http.ResponseWriter, request *http.Request) {
			token, err := ioutil.ReadAll(request.Body)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				_, _ = writer.Write([]byte{})
				panic(err)
			}
			waitForToken <- string(token)
		})
		listener, err := net.Listen("tcp", ":5364")
		if err != nil {
			panic(err)
		}
		defer listener.Close()
		log.Println("http server started")
		go func() {
			http.Serve(listener, mux)
		}()
		go func() {
			log.Println("opening twitch login in browser")
			if err := open.Browse(twitch.AuthorizeUrl("http://localhost:5364/twitch_login.html")); err != nil {
				panic(err)
			}
		}()
		log.Println("waiting for token")
		token := <- waitForToken
		log.Println("found token")
		log.Println("saving...")
		integration := twitch.LoadTwitch()
		integration.Token = token
		log.Println("testing connection...")
		err = integration.Connect()
		if err != nil {
			log.Printf("failed to authenticate with twitch: %s", err)
			os.Exit(1)
			return
		}
		integration.Save()
	},
}
