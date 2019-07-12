package controllers

import (
	"strings"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Me() revel.Result {
	return c.Render()
}

func (c App) Level5() revel.Result {
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

func (c App) Blog() revel.Result {
	return c.Render()
}

var allRobots [5]string

func (c App) Robot() revel.Result {
	allRobots := [5]string{"sandstorm", "sideslip", "faultline", "echo", "blackout"}

	robot := strings.ToLower(c.Params.Route.Get("robot"))
	var valid = false

	for i := 0; i < 5; i++ {
		if robot == allRobots[i] {
			valid = true
		}

	}

	if !valid {
		robot = "Invalid Robot " + robot + "!"
	}

	robot = strings.Title(robot)

	return c.Render(robot)
}
