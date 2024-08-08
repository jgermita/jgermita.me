package controllers

import (
	"strings"

	"github.com/jgermita/jgermita.me/app/model"

	"github.com/revel/revel"
	"io/ioutil"
	"net/http"

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

func (c App) Belts() revel.Result {
	return c.Render()
}

func (c App) Chain() revel.Result {
	return c.Render()
}

func (c App) Projects() revel.Result {
	return c.Render()
}

func (c App) First() revel.Result {

	db := new(model.Database)

	var bots = db.GetFirstRobots()

	return c.Render(bots)
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
func (c App) Card() revel.Result {
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

func (c App) Crresources() revel.Result {
	return c.Render()
}

func (c App) Level5() revel.Result {

	

	db := new(model.Database)

	var bots = db.GetAllRobots()
	var events = db.GetAllEvents()
	var upcoming = db.GetUpcomingEvents()
	var record = db.GetRecord()
	var videos = db.GetHighlights()

	return c.Render(bots, events, record, videos, upcoming)
}

func (c App) Robot() revel.Result {

	robot := strings.ToLower(c.Params.Route.Get("robot"))
	robot = strings.Title(robot)
	db := new(model.Database)

	var bot = db.GetRobot(robot)

	rank := "--"
	if(bot.Rce != "") {
		resp, err := http.Get(bot.Rce)
		if err != nil {
			println(err)
		}
		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			println(err)
		}

		rankFound := false
		
		for _, element := range strings.Split(string(body), "\n") {
			if(rankFound) {
				//rank 
				rank = (strings.Trim(element, " "));
				break
			}
			if(strings.Contains(element, "<div class=\"resource-header-rank \">")) {
				rankFound = true
			}
		}
	} else {
	}
	
	return c.Render(bot, rank)
}

func (c App) Fight() revel.Result {

	id := strings.ToLower(c.Params.Route.Get("id"))

	db := new(model.Database)

	var fight = db.GetFight(id)
	return c.Render(fight)
}
