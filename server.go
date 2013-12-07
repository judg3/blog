package main

import (
	"blog/controllers"
	"blog/forms"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/post/:id", controllers.GetPost)
	m.Get("/posts", controllers.ListPosts)
	m.Post("/post", binding.Form(forms.Post{}), controllers.CreatePost)

	m.Run()
}
