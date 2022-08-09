package repository

import (
	"MIET-TelegramBot/internal/app/store/models"
	"fmt"
	"log"
)

type UserRepository struct {
	repository *Repository
}

func (u *UserRepository) Create(user *models.UserModel) error {

	_, err := u.repository.db.Exec(
		`INSERT INTO miet_tgbot_db.member (
			user_tg_id,
			first_name,
			last_name,
			username,
			group_name
		) 
		VALUES (
			$1,
			$2,
			$3
			$4,
			$5
		) `,
		user.UserTgId,
		user.FisrtName,
		user.LastName,
		user.UserName,
		user.Group,
	)

	if err != nil {
		log.Println(fmt.Sprintln("Error: write in database failed, %s", err.Error()))
		return err
	}

	return nil
}

func (u *UserRepository) Delete(user *models.UserModel) error {
	_, err := u.repository.db.Exec(`
		DELETE FROM miet_tgbot_db.member 
		WHERE user_tg_id = $1`,
		user.UserTgId)

	if err != nil {
		log.Println(fmt.Sprintln("Error: remove from database failed, %s", err.Error()))
		return err
	}
	return nil
}

func (u *UserRepository) Contains(user *models.UserModel) (bool, error) {
	var count bool
	row := u.repository.db.QueryRow(`
		SELECT COUNT(*)
		WHERE  user_tg_id = $1 
			AND first_name = $2
			AND last_name = $3 
			AND username = $4
	`, user.UserTgId, user.FisrtName, user.LastName, user.UserName)

	if err := row.Scan(&count); err != nil {
		log.Println(fmt.Sprintln("Error: get group from database failed, %s", err.Error()))
		return false, err
	}
	return bool(count), nil
}

func (u *UserRepository) Group(user *models.UserModel) (string, error) {
	var group string
	row := u.repository.db.QueryRow(`
		SELECT group_name
		WHERE  user_tg_id = $1 
			AND first_name = $2
			AND last_name = $3 
			AND username = $4
	`, user.UserTgId, user.FisrtName, user.LastName, user.UserName)

	if err := row.Scan(&group); err != nil {
		log.Println(fmt.Sprintln("Error: get group from database failed, %s", err.Error()))
		return "", err
	}

	return group, nil
}
