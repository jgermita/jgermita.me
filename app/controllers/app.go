package controllers

import (
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
