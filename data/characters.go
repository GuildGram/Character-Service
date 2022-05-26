package data

import (
	"encoding/json"
	"fmt"
	"io"
)

type GuildData struct {
	GuildID   string `json:"guildid"`
	GuildRole string `json:"guildrole"`
}

type Character struct {
	UserID           string `json:"userid"`
	Class            string `json:"class"`
	CharaterName     string `json:"name"`
	RegionServerName string `json:"regionserver"`
	CharacterLevel   int    `json:"characterlevel"`
	RosterLevel      int    `json:"rosterevel"`
	Ilvl             int    `json:"ilvl"`
	GuildID          string `json:"guildid"`
	GuildRole        string `json:"guildrole"`
}

func (c *Character) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}

func (c *GuildData) FromJSON(r io.Reader) error {
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

func UpdateCharacterGuild(id string, gId string, role string) error {
	char, _, err := findChar(id)
	if err != nil {
		return err
	}
	char.GuildID = gId
	char.GuildRole = role
	return err
}

func UpdateCharacter(id string, c *Character) error {
	_, pos, err := findChar(id)
	if err != nil {
		return err
	}

	c.UserID = id
	characterList[pos] = c
	return err
}

var ErrCharNotFound = fmt.Errorf("char Not found")

func findChar(id string) (*Character, int, error) {
	for i, c := range characterList {
		if c.UserID == id {
			return c, i, nil
		}
	}
	return nil, -1, ErrCharNotFound
}

func AddCharacter(c *Character) {
	characterList = append(characterList, c)
}

func DeleteCharacter(id string) error {
	_, pos, err := findChar(id)
	if err != nil {
		return err
	}
	characterList[pos] = characterList[len(characterList)-1]
	characterList[len(characterList)-1] = nil
	characterList = characterList[:len(characterList)-1]
	return err
}

func GetCharacter(id string) (*Character, error) {
	_, pos, err := findChar(id)
	if err != nil {
		return nil, err
	}
	return characterList[pos], err
}

func GetCharactersByGuild(id string) ([]*Character, error) {
	var cList []*Character
	for _, c := range characterList {
		if c.GuildID == id {
			cList = append(cList, c)
		}
	}
	if cList != nil {
		return cList, nil
	}
	return nil, ErrCharNotFound
}

var characterList = []*Character{
	{
		UserID:           "1",
		CharaterName:     "Nemoi",
		Class:            "Striker",
		RegionServerName: "EUC-Sceptrum",
		CharacterLevel:   53,
		RosterLevel:      68,
		Ilvl:             1355,
		GuildID:          "G1",
		GuildRole:        "Owner",
	},
	{
		UserID:           "2",
		CharaterName:     "Mjc",
		Class:            "Berserk",
		RegionServerName: "EUC-Sceptrum",
		CharacterLevel:   53,
		RosterLevel:      60,
		Ilvl:             1368,
		GuildID:          "G2",
		GuildRole:        "Owner",
	},
	{
		UserID:           "3",
		CharaterName:     "Leopewpew",
		Class:            "Archer",
		RegionServerName: "EUC-Sceptrum",
		CharacterLevel:   53,
		RosterLevel:      60,
		Ilvl:             1368,
		GuildID:          "G1",
		GuildRole:        "Member",
	},
}
