package main

import (
	"net/http"

	"github.com/GuildGram/Character-Service/handlers"
)

func main() {

	ch := handlers.NewCreate()
	hh := handlers.NewCreate()
	println(ch)
	http.ListenAndServe(":9090", nil)
}
