package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Character struct {
	ID               int    `json:"id"`
	CharaterName     string `json:"name"`
	RegionServerName string `json:"region-server"`
	CharacterLevel   int    `json:"characterlevel"`
	RosterLevel      int    `json:"rosterLevel"`
	GuildName        string `json:"guildName"`
	GuildRole        string `json:"guildRole"`
	//Might delete for internal use for now
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

func (c *Character) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}

type Characters []*Character

func (c *Characters) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func (c *Character) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func GetCharacters() Characters {
	return characterList
}

func UpdateCharacter(id int, c *Character) error {
	_, pos, err := findChar(id)
	if err != nil {
		return err
	}

	c.ID = id
	characterList[pos] = c
	return err
}

var ErrCharNotFound = fmt.Errorf("Char Not found")

func findChar(id int) (*Character, int, error) {
	for i, c := range characterList {
		if c.ID == id {
			return c, i, nil
		}
	}
	return nil, -1, ErrCharNotFound
}

func AddCharacter(c *Character) {
	c.ID = GetNextID()
	characterList = append(characterList, c)
}

func GetNextID() int {
	return characterList[len(characterList)-1].ID + 1
}

func DeleteCharacter(id int) error {
	_, pos, err := findChar(id)
	if err != nil {
		return err
	}
	characterList[pos] = characterList[len(characterList)-1]
	characterList[len(characterList)-1] = nil
	characterList = characterList[:len(characterList)-1]
	return err
}

func GetCharacter(id int) (*Character, error) {
	_, pos, err := findChar(id)
	if err != nil {
		return nil, err
	}
	return characterList[pos], err
}

var characterList = []*Character{
	&Character{
		ID:               1,
		CharaterName:     "Nemoi",
		RegionServerName: "EUC-Sceptrum",
		CharacterLevel:   53,
		RosterLevel:      68,
		GuildName:        "FontysICT",
		GuildRole:        "Owner",
		CreatedOn:        time.Now().UTC().String(),
		UpdatedOn:        time.Now().UTC().String(),
	},
	&Character{
		ID:               2,
		CharaterName:     "Mjc",
		RegionServerName: "EUC-Sceptrum",
		CharacterLevel:   53,
		RosterLevel:      60,
		GuildName:        "InternsGuild",
		GuildRole:        "Member",
		CreatedOn:        time.Now().UTC().String(),
		UpdatedOn:        time.Now().UTC().String(),
	},
}
