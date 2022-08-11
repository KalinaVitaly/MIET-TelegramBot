package models

import (
	"fmt"
	"log"
)

type UserModel struct {
	UserTgId  int64
	FisrtName string
	LastName  string
	UserName  string
	Group     string
}

func CreateUserModel(id int64, firstName, lastName, userName, group string) *UserModel {
	return &UserModel{
		UserTgId:  id,
		FisrtName: firstName,
		LastName:  lastName,
		UserName:  userName,
		Group:     group,
	}
}

func (u *UserModel) SetGroup(group string) {
	u.Group = group
}

func (u *UserModel) ValidGroup(group string) bool {
	for _, _group := range groupsList {
		log.Println(fmt.Sprintln("Groups : %s : %s", _group, group))
		if _group == group {
			return true
		}
	}
	return false
}
