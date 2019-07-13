package model

type Fight struct {
	id            int
	Event         string
	MyRobot       string
	OpponentRobot string
	MyWin         bool
	Video         string
	VideoExists   bool
}
