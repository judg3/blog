package controllers

import (
	"github.com/codegangsta/martini"
)

func Hello(params martini.Params) string {
	return "Hello " + params["name"]
}
