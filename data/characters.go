package data

import (
	"encoding/json"
	"io"
	"time"
)

type Character struct {
	CharaterName     string `json:"name"`
	RegionServerName string `json:"region-server"`
	ItemLevel        int    `json:"ilvl"`
	CharacterLevel   int    `json:"characterlevel"`
	RosterLevel      int    `json:"rosterLevel"`
	GuildName        string `json:"guildName"`
	GuildRole        string `json:"guildRole"`
	//Might delete for internal use for now
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

type Characters []*Character

func (c *Characters) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func GetCharacters() Characters {
	return characterList
}

var characterList = []*Character{
	&Character{
		CharaterName:     "Nemoi",
		RegionServerName: "EUC-Sceptrum",
		ItemLevel:        1355,
		CharacterLevel:   53,
		RosterLevel:      68,
		GuildName:        "FontysICT",
		GuildRole:        "Owner",
		CreatedOn:        time.Now().UTC().String(),
		UpdatedOn:        time.Now().UTC().String(),
	},
	&Character{
		CharaterName:     "Mjc",
		RegionServerName: "EUC-Sceptrum",
		ItemLevel:        1340,
		CharacterLevel:   53,
		RosterLevel:      60,
		GuildName:        "InternsGuild",
		GuildRole:        "Member",
		CreatedOn:        time.Now().UTC().String(),
		UpdatedOn:        time.Now().UTC().String(),
	},
}
