package model

type Robot struct {
	Id          int
	Name        string
	WeightClass string
	Status      string
	Wins        int
	Losses      int
	Events      []Event
	Fights      []Fight
	Logo        string
	Img         string
	Media       []string
	Weapon      string
	Desc        string
	Rgb         string

	Winrate float64
}
