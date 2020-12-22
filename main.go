package main

import (
	"atwitchant/pkg/middleware"
	"atwitchant/pkg/twitch"
	"log"
	"net/http"
)

func main() {
	log.Println("welcome to atwitchant")
	twitchIntegration := &twitch.Integration{}
	loggerMiddleware := &middleware.Logger{}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.Handle("/auth/twitch", twitchIntegration.Authorize())
	handler := loggerMiddleware.Log(mux)

	log.Println("staring http server on http://localhost:5364")
	if err := http.ListenAndServe(":5364", handler); err != http.ErrServerClosed {
		panic(err)
	}
}
