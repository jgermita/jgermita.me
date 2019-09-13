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

func (c App) About() revel.Result {
	return c.Render()
}

func (c App) Dashboard() revel.Result {
	team := strings.ToLower(c.Params.Route.Get("team"))

	return c.Render(team)
}
