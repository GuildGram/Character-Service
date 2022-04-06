package handlers

import (
	"log"
	"net/http"

	"github.com/GuildGram/Character-Service.git/data"
)

type Character struct {
	l *log.Logger
}

func NewCharacter(l *log.Logger) *Character {
	return &Character{l}
}

func (c *Character) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.GetCharacters(rw, r)
		return
	}

	//if not get then throw error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (c *Character) GetCharacters(rw http.ResponseWriter, h *http.Request) {
	listChars := data.GetCharacters()
	err := listChars.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to Marshal Json", http.StatusInternalServerError)
	}
}
