package store

import "./models"

type UserRep interface {
	Create(*models.UserModel) error
	Delete(*models.UserModel) error
	Contains(user *models.UserModel, userDeauth bool) (bool, error)
	Group(*models.UserModel) (string, error)
}
