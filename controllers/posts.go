package controllers

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/judg3/blog/forms"
	"github.com/judg3/blog/models"
	"github.com/judg3/blog/services"
)

func GetPost(params martini.Params, r render.Render) {
	hd := services.GetDb()

	var results []models.Posts
	hd.Where("id", "=", params["id"]).Limit(1).Find(&results)

	r.JSON(200, results)
}

func ListPosts(params martini.Params, r render.Render) {
	hd := services.GetDb()

	var results []models.Posts
	hd.Find(&results)

	r.JSON(200, results)
}

func CreatePost(post forms.Post, params martini.Params, formErr binding.Errors, r render.Render) {

	hd := services.GetDb()
	tx := hd.Begin()

	newpost := models.Posts{Title: post.Title, Body: post.Content}
	id, err := tx.Save(&newpost)

	if err != nil {
		panic(err)
	}

	err = tx.Commit()

	r.JSON(200, id)
}
