package controllers

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/judg3/blog/forms"
	"github.com/judg3/blog/services"
	"log"
)

func Login(login forms.Login, params martini.Params, r render.Render, l *log.Logger) {
	result := services.GetUserService().Login(login.Username, login.Password)
	r.JSON(200, result)
}
