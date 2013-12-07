package services

import (
	"github.com/eaigner/hood"
)

var instantiated *hood.Hood
var err error

func GetDb() *hood.Hood {

	if instantiated == nil {

		instantiated, err = hood.Load("db/config.json", "development")
		if err != nil {
			panic(err)
		}
	}
	return instantiated
}
