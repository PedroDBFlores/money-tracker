package service

import "money-tracker/api/internal/user"

type CreateUser struct {
	repo user.UserRepository
}

func (c *CreateUser) Execute(user user.User) (int64, error) {
	return c.repo.Save(user)
}
