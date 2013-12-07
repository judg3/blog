package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/judg3/blog/controllers"
	"github.com/judg3/blog/forms"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/post/:id", controllers.GetPost)
	m.Get("/posts", controllers.ListPosts)
	m.Post("/post", binding.Form(forms.Post{}), controllers.CreatePost)

	m.Run()
}
