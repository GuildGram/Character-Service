package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if r.Method == http.MethodPost {
		c.addCharacter(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		//check if only one ID was retrieved
		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		//ceck if capture groups created succesfully
		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			return
		}

		c.updateCharacters(id, rw, r)
		return
	}

	//if not get then throw error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
func (c *Character) updateCharacters(id int, rw http.ResponseWriter, r *http.Request) {
	c.l.Println("HANDLE PUT CHARACTER")

	char := &data.Character{}
	err := char.FromJSON(r.Body)
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

func (c *Character) addCharacter(rw http.ResponseWriter, r *http.Request) {
	c.l.Println("HANDLE POST CHARACTER")

	char := &data.Character{}
	err := char.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddCharacter(char)
}
