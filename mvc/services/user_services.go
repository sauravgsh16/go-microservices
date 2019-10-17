package services

import (
	"github.com/sauravgsh16/microservices/mvc/domain"
	"github.com/sauravgsh16/microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.AppError) {
	return domain.GetUser(userId)
}
