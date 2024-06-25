package models

import (
	"strconv"
)

type User struct {
	Id     int
	Name   string
	Email  string
	Status bool
}

func (u *User) SetUser(id int, name, email string, status bool) {
	u.Id = id
	u.Name = name
	u.Email = email
	u.Status = status
}

func (u *User) ToString() string {
	return "Id: " + strconv.Itoa(u.Id) + " Name: " + u.Name + " Email: " + u.Email + " Status: " + strconv.FormatBool(u.Status)
}
