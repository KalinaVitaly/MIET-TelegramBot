package store

import "MIET-TelegramBot/internal/app/store/models"

type UserRep interface {
	Create(*models.UserModel) error
	Delete(*models.UserModel) error
	Contains(*models.UserModel) error
	Group(*models.UserModel) (string, error)
}
