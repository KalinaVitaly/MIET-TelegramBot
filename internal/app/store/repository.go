package store

import "./models"

type UserRep interface {
	Create(*models.UserModel) error
	Delete(*models.UserModel) error
	Contains(*models.UserModel) (bool, error)
	Group(*models.UserModel) (string, error)
}
