package models

import (
	"fmt"
	"log"
	"strings"
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

func (u *UserModel) ValidGroup(_group string) (string, string, bool) {
	for groupEn, groupRus := range groupsList {
		log.Println(fmt.Sprintln("Groups : %s : %s", _group, groupEn))
		if groupEn == strings.ToUpper(_group) || groupRus == strings.ToUpper(_group) {
			return groupEn, groupRus, true
		}
	}
	return "", "", false
}
