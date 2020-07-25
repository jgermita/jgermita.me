package model

type Fight struct {
	Id            int
	Event         string
	MyRobot       string
	OpponentRobot string
	MyWin         bool
	Video         string
	VideoExists   bool
	Recap         string
	IsRumble      bool
	Type          string
}
