package main

import (
	"github.com/eaigner/hood"
)

func (m *M) AddUsernameIndex_1385792598_Up(hd *hood.Hood) {
	hd.CreateIndex("users", "name_index", true, "first", "last")
}

func (m *M) AddUsernameIndex_1385792598_Down(hd *hood.Hood) {
	// TODO: implement
}
