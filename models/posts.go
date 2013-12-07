package models

import (
	"github.com/eaigner/hood"
)

type Posts struct {
	Id        hood.Id
	Title     string `validate:"presence len(2:255)" sql:"size(255)"`
	Body      string `validate:"presence"`
	AuthorId  int64
	CreatedAt hood.Created
	UpdatedAt hood.Updated
}
