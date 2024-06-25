package repository

import (
	"github.com/chava.gnolasco/godesde0/models"
)

func CreateUser(id int, name string, email string, status bool) models.User {
	user := new(models.User)
	user.SetUser(id, name, email, status)
	return *user
}
