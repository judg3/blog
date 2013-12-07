package main

import (
	"github.com/eaigner/hood"
)

func (m *M) CreatePostsTable_1385803004_Up(hd *hood.Hood) {
	type Posts struct {
		Id        hood.Id
		Title     string `validate:"presence len(2:255)" sql:"size(255)"`
		Body      string `validate:"presence"`
		AuthorId  int64
		CreatedAt hood.Created
		UpdatedAt hood.Updated
	}
	hd.CreateTable(&Posts{})
}

func (m *M) CreatePostsTable_1385803004_Down(hd *hood.Hood) {
	// TODO: implement
}
