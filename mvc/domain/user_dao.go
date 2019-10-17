package domain

import (
	"fmt"
	"net/http"

	"github.com/sauravgsh16/microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: &User{
			ID:        123,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@test.com",
		},
	}
)

func GetUser(userId int64) (*User, *utils.AppError) {
	user, found := users[userId]
	if !found {
		return nil, &utils.AppError{
			ID:      http.StatusNotFound,
			Message: fmt.Sprintf("User: %d not found", userId),
		}
	}
	return user, nil
}
