package main

import (
	"github.com/eaigner/hood"
)

func (m *M) CreateUsersTable_1385792513_Up(hd *hood.Hood) {
	type Users struct {
		Id    hood.Id
		First string
		Last  string
	}
	hd.CreateTable(&Users{})
}

func (m *M) CreateUsersTable_1385792513_Down(hd *hood.Hood) {
	// TODO: implement
}
