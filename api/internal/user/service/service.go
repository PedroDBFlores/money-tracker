package service

import "money-tracker/api/internal/user"

type UserService struct {
	CreateUser CreateUser
	GetByEmail GetByEmail
}

func NewUserService(repo user.UserRepository) UserService {
	return UserService{
		CreateUser: CreateUser{repo: repo},
		GetByEmail: GetByEmail{repo: repo},
	}
}
