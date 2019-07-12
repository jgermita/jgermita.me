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

func (c App) Robot() revel.Result {
	var allRobots = [5]string{"sandstorm", "sideslip", "faultline", "echo", "blackout"}

	robot := strings.ToLower(c.Params.Route.Get("robot"))
	var valid = false

	var img string

	for i := 0; i < 5; i++ {
		if robot == allRobots[i] {
			valid = true
		}

	}

	if !valid {
		robot = "Invalid Robot " + robot + "!"
		img = "l5logo.png"
	} else {
		robot = strings.Title(robot)
		img = strings.ToLower(robot) + ".png"
	}

	img = "/public/img/level5assets/" + img

	return c.Render(robot, img)
}
