package controllers

import (
	"math/rand"
	"strings"
	"time"

	"github.com/jgermita/jgermita.me/app/model"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller

	Database *model.Database
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Me() revel.Result {
	return c.Render()
}

func (c App) Projects() revel.Result {
	return c.Render()
}

func (c App) First() revel.Result {
	return c.Render()
}

func (c App) Software() revel.Result {
	return c.Render()
}
func (c App) Countdown() revel.Result {
	return c.Render()
}

func (c App) Countdown5012() revel.Result {
	return c.Render()
}

func (c App) Blog() revel.Result {
	return c.Render()
}

func (c App) Dt() revel.Result {
	return c.Render()
}
func (c App) Dtbrushed() revel.Result {
	return c.Render()
}

func (c App) W() revel.Result {
	return c.Render()
}

func (c App) Todo() revel.Result {
	return c.Render()
}
func (c App) Services() revel.Result {
	return c.Render()
}

func CreateUser(name string, ign string) model.User {
	var thisUser model.User

	thisUser.Name = name
	thisUser.Ign = ign

	return thisUser
}

func (c App) League() revel.Result {

	users := []model.User{}
	users = append(users, CreateUser("Jeremy", "King Cheesecake"))
	users = append(users, CreateUser("CP", "chrisr2"))
	users = append(users, CreateUser("Brian", "CptBubbleFace"))
	users = append(users, CreateUser("Joseph", "Danklydankmemes"))
	users = append(users, CreateUser("Angel", "DeathAngel101"))
	users = append(users, CreateUser("Javi", "Midlane Hokage"))
	users = append(users, CreateUser("John", "TastiestMemes"))
	users = append(users, CreateUser("Marcus", "Toaster130"))
	users = append(users, CreateUser("Justin", "BustinJustin420"))
	users = append(users, CreateUser("Mark", "NotMark420"))
	users = append(users, CreateUser("Joaquin", "NotJoaquin420"))
	users = append(users, CreateUser("Andy", ""))
	users = append(users, CreateUser("Este", ""))

	p1 := strings.ToLower(c.Params.Values.Get("p1"))
	p2 := strings.ToLower(c.Params.Values.Get("p2"))
	p3 := strings.ToLower(c.Params.Values.Get("p3"))
	p4 := strings.ToLower(c.Params.Values.Get("p4"))
	p5 := strings.ToLower(c.Params.Values.Get("p5"))
	p6 := strings.ToLower(c.Params.Values.Get("p6"))
	p7 := strings.ToLower(c.Params.Values.Get("p7"))
	p8 := strings.ToLower(c.Params.Values.Get("p8"))
	p9 := strings.ToLower(c.Params.Values.Get("p9"))
	p10 := strings.ToLower(c.Params.Values.Get("p10"))

	allRand := strings.ToLower(c.Params.Values.Get("allrand"))
	showRoles := strings.ToLower(c.Params.Values.Get("showroles"))

	players := []string{}

	players = append(players, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10)

	red := []string{}
	blue := []string{}

	if allRand == "allrand" {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(players), func(i, j int) { players[i], players[j] = players[j], players[i] })
	}

	red = append(red, players[0], players[1], players[2], players[3], players[4])
	blue = append(blue, players[5], players[6], players[7], players[8], players[9])

	return c.Render(red, blue, showRoles, users)
}

func (c App) Level5() revel.Result {

	db := new(model.Database)

	var bots = db.GetAllRobots()
	var events = db.GetAllEvents()
	var record = db.GetRecord()
	var videos = db.GetHighlights()

	return c.Render(bots, events, record, videos)
}
func (c App) Robot() revel.Result {

	robot := strings.ToLower(c.Params.Route.Get("robot"))
	robot = strings.Title(robot)
	db := new(model.Database)

	var bot = db.GetRobot(robot)
	return c.Render(bot)
}
