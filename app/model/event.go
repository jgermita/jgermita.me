package model

type Event struct {
	id          int
	Name        string
	Finish      string
	Date        string
	Report      string
	Robot       string
	Video       string
	VideoExists bool
	WinCount    int
	LossCount   int
	Website     string
	Location    string
	Bracket     string

	Fights []Fight
}
