package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Character struct {
	UserID           int    `json:"userid"`
	Class            string `json:"class"`
	CharaterName     string `json:"name"`
	RegionServerName string `json:"region-server"`
	CharacterLevel   int    `json:"characterlevel"`
	RosterLevel      int    `json:"rosterLevel"`
	Ilvl             int    `json:"ilvl"`
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

	c.UserID = id
	characterList[pos] = c
	return err
}

var ErrCharNotFound = fmt.Errorf("Char Not found")

func findChar(id int) (*Character, int, error) {
	for i, c := range characterList {
		if c.UserID == id {
			return c, i, nil
		}
	}
	return nil, -1, ErrCharNotFound
}

func AddCharacter(c *Character) {
	c.UserID = GetNextID()
	characterList = append(characterList, c)
}

func GetNextID() int {
	return characterList[len(characterList)-1].UserID + 1
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
		UserID:           1,
		CharaterName:     "Nemoi",
		Class:            "Striker",
		RegionServerName: "EUC-Sceptrum",
		CharacterLevel:   53,
		RosterLevel:      68,
		Ilvl:             1355,
		GuildName:        "FontysICT",
		GuildRole:        "Owner",
		CreatedOn:        time.Now().UTC().String(),
		UpdatedOn:        time.Now().UTC().String(),
	},
	&Character{
		UserID:           2,
		CharaterName:     "Mjc",
		Class:            "Berserk",
		RegionServerName: "EUC-Sceptrum",
		CharacterLevel:   53,
		RosterLevel:      60,
		Ilvl:             1368,
		GuildName:        "InternsGuild",
		GuildRole:        "Member",
		CreatedOn:        time.Now().UTC().String(),
		UpdatedOn:        time.Now().UTC().String(),
	},
}
