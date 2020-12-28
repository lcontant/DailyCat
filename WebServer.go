package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var discordHandler DiscordHandler
var configuration Configuration
func handleOauthReturn(w http.ResponseWriter, r *http.Request) {
	exchangeCodeResponse := discordHandler.exchange_code(r.URL.Query().Get("code"))
	resp, err := json.Marshal(exchangeCodeResponse)
	if err != nil {
		fmt.Println(err)
	}
	code, err := w.Write(resp)
	if err != nil {
		fmt.Println(err)
		fmt.Println(code)
	}
}
func main() {
	env := ""
	if len(os.Args) >= 2 {
		env = os.Args[1]
	}
	filename := "config.json"
	if len(env) > 0 {
		filename = "config." + env + ".json"
	}
	configuration = newConfiguration(filename)
	discordHandler = DiscordHandler{configuration: configuration}
	http.HandleFunc("/discordRedirect", handleOauthReturn)
	http.Handle("/", http.StripPrefix("", http.FileServer(http.Dir(""))))
	log.Fatal(http.ListenAndServeTLS(":4444", "server.crt", "server.key", nil))
}
