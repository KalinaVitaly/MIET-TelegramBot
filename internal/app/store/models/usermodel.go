package models

type UserModel struct {
	UserTgId  uint64
	FisrtName string
	LastName  string
	UserName  string
	Group     string
}

func CreateUserModel(id uint64, firstName, lastName, userName, group string) *UserModel {
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
