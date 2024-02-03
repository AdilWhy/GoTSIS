package models

import "errors"

type Companion struct {
	Name       string `json:"name"`
	Race       string `json:"race"`
	Class      string `json:"class"`
	Background string `json:"background"`
	Attributes struct {
		Str int `json:"Str"`
		Dex int `json:"Dex"`
		Con int `json:"Con"`
		Int int `json:"Int"`
		Wis int `json:"Wis"`
		Cha int `json:"Cha"`
	} `json:"attributes"`
}

var Companions = []Companion{
	{
		Name:       "Astarion",
		Race:       "High Elf",
		Class:      "Rogue",
		Background: "Charlatan",
		Attributes: struct {
			Str int `json:"Str"`
			Dex int `json:"Dex"`
			Con int `json:"Con"`
			Int int `json:"Int"`
			Wis int `json:"Wis"`
			Cha int `json:"Cha"`
		}{8, 17, 14, 13, 13, 10},
	},
	{
		Name:       "Shadowheart",
		Race:       "High Half-Elf",
		Class:      "Cleric",
		Background: "Acolyte",
		Attributes: struct {
			Str int `json:"Str"`
			Dex int `json:"Dex"`
			Con int `json:"Con"`
			Int int `json:"Int"`
			Wis int `json:"Wis"`
			Cha int `json:"Cha"`
		}{13, 13, 14, 10, 17, 8},
	},
	{
		Name:       "Gale",
		Race:       "Human",
		Class:      "Wizard",
		Background: "Sage",
		Attributes: struct {
			Str int `json:"Str"`
			Dex int `json:"Dex"`
			Con int `json:"Con"`
			Int int `json:"Int"`
			Wis int `json:"Wis"`
			Cha int `json:"Cha"`
		}{8, 13, 15, 17, 10, 12},
	},
	{
		Name:       "Lae'zel",
		Race:       "Githyanki",
		Class:      "Fighter",
		Background: "Soldier",
		Attributes: struct {
			Str int `json:"Str"`
			Dex int `json:"Dex"`
			Con int `json:"Con"`
			Int int `json:"Int"`
			Wis int `json:"Wis"`
			Cha int `json:"Cha"`
		}{17, 13, 15, 10, 12, 8},
	},
	{
		Name:       "Karlach",
		Race:       "Zariel Tiefling",
		Class:      "Barbarian",
		Background: "Outlander",
		Attributes: struct {
			Str int `json:"Str"`
			Dex int `json:"Dex"`
			Con int `json:"Con"`
			Int int `json:"Int"`
			Wis int `json:"Wis"`
			Cha int `json:"Cha"`
		}{17, 13, 15, 8, 12, 10},
	},
	{
		Name:       "Wyll",
		Race:       "Human",
		Class:      "Warlock",
		Background: "Folk Hero",
		Attributes: struct {
			Str int `json:"Str"`
			Dex int `json:"Dex"`
			Con int `json:"Con"`
			Int int `json:"Int"`
			Wis int `json:"Wis"`
			Cha int `json:"Cha"`
		}{8, 13, 14, 13, 10, 17},
	},
}

func GetCompanion(name string) (*Companion, error) {
	for _, c := range Companions {
		if c.Name == name {
			return &c, nil
		}
	}
	return nil, errors.New("no such companion")
}
