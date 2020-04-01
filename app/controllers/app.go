package controllers

import (
	"strings"

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

func (c App) W() revel.Result {
	return c.Render()
}

func (c App) Todo() revel.Result {
	return c.Render()
}

func (c App) Level5() revel.Result {

	db := new(model.Database)

	var bots = db.GetAllRobots()
	var events = db.GetAllEvents()
	var record = db.GetRecord()

	return c.Render(bots, events, record)
}
func (c App) Robot() revel.Result {

	robot := strings.ToLower(c.Params.Route.Get("robot"))
	robot = strings.Title(robot)
	db := new(model.Database)

	var bot = db.GetRobot(robot)
	return c.Render(bot)
}
