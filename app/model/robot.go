package model

type Robot struct {
	Id          int
	Name        string
	WeightClass string
	Status      string
	Wins        int
	Losses      int
	Events      string
	Fights      []Fight
	Logo        string
	Img         string
	Media       []string
	Weapon      string

	Winrate float64
}
