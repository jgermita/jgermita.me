package model

type User struct {
	Name string
	Ign  string

	TopRating float64
	JngRating float64
	MidRating float64
	BotRating float64
	SupRating float64

	Winrate float64
}
