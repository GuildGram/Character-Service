package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/GuildGram/Character-Service.git/data"
	"github.com/gorilla/mux"
)

type Character struct {
	l *log.Logger
}

func NewCharacter(l *log.Logger) *Character {
	return &Character{l}
}

func (c *Character) UpdateCharacters(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert ID", http.StatusBadRequest)
		return
	}

	c.l.Println("HANDLE PUT CHARACTER", id)

	char := &data.Character{}
	err = char.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateCharacter(id, char)
	if err == data.ErrCharNotFound {
		http.Error(rw, "Char not found", http.StatusNotFound)
	}
	if err != nil {
		http.Error(rw, "Char not found", http.StatusInternalServerError)
	}
}

func (c *Character) GetCharacters(rw http.ResponseWriter, h *http.Request) {
	c.l.Println("HANDLE GET CHARACTERS")
	listChars := data.GetCharacters()
	err := listChars.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to Marshal Json", http.StatusInternalServerError)
	}
}

func (c *Character) AddCharacter(rw http.ResponseWriter, r *http.Request) {
	c.l.Println("HANDLE POST CHARACTER")

	char := &data.Character{}
	err := char.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddCharacter(char)
}

func (c *Character) DeleteCharacter(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert ID", http.StatusBadRequest)
		return
	}

	c.l.Println("HANDLE DELETE CHARACTER", id)

	err = data.DeleteCharacter(id)
	if err != nil {
		http.Error(rw, "Char not found", http.StatusInternalServerError)
	}
}

func (c *Character) GetCharacter(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert ID", http.StatusBadRequest)
		return
	}

	c.l.Println("HANDLE GET 1 CHARACTER", id)

	char, err2 := data.GetCharacter(id)
	if err2 != nil {
		http.Error(rw, "Char not found", http.StatusInternalServerError)
	}
	char.ToJSON(rw)
}
