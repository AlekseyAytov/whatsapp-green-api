package main

import (
	"net/http"

	"github.com/AlekseyAytov/whatsapp-green-api/internal"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc(`/`, http.FileServer(http.Dir("./web/template/")).ServeHTTP)
	mux.HandleFunc(`/settings`, internal.SettingsHandler)
	mux.HandleFunc(`/state`, internal.StateHandler)
	mux.HandleFunc(`/send`, internal.SendMessage)
	mux.HandleFunc(`/file`, internal.SendFile)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
