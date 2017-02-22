package controllers

import (
	"github.com/revel/revel"
	"fmt"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	fmt.Printf("test")
	return c.Render()
}
