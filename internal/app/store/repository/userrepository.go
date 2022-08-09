package repository

import "MIET-TelegramBot/internal/app/store/models"

type UserRepository struct {
	repository *Repository
}

func (u *UserRepository) Create(*models.UserModel) error {

	return nil
}

func (u *UserRepository) Delete(*models.UserModel) error {

	return nil
}

func (u *UserRepository) Contains(*models.UserModel) error {

	return nil
}

func (u *UserRepository) Group(*models.UserModel) (string, error) {

	return "", nil
}
