package main

import (
	"log"
	"net/http"
	"os"

	"github.com/GuildGram/Character-Service.git/handlers"
)

func main() {
	l := log.New(os.Stdout, "character-api", log.LstdFlags)
	ch := handlers.NewCreate(l)

	sm := http.NewServeMux()
	sm.Handle("/", ch)

	// http.HandleFunc()
	http.ListenAndServe(":9090", nil)
}
