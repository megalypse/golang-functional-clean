package user

import "github.com/megalypse/golang-functional-clean/internal/domain/models"

type UserDbe struct {
	Username string `bson:"username"`
	Age      int    `bson:"age"`
}

func (e UserDbe) ToEntity() models.User {
	return models.User{
		Username: e.Username,
		Age:      e.Age,
	}
}
