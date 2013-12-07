package db

import (
	"github.com/eaigner/hood"
)

type Users struct {
	Id    hood.Id
	First string
	Last  string
}

func (table *Users) Indexes(indexes *hood.Indexes) {
	indexes.AddUnique("name_index", "first", "last")
}

type Posts struct {
	Id        hood.Id
	Title     string `validate:"presence len(2:255)" sql:"size(255)"`
	Body      string `validate:"presence"`
	AuthorId  int64
	CreatedAt hood.Created
	UpdatedAt hood.Updated
}
