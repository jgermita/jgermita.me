package model

type Robot struct {
	Id             int
	Name           string
	Name2          string
	WeightClass    string
	Status         string
	Wins           int
	Losses         int
	Events         []Event
	Fights         []Fight
	Logo           string
	Img            string
	Media          string
	Weapon         string
	Desc           string
	Rgb            string
	Versions       []string
	CurrentVersion string

	Winrate float64
}
