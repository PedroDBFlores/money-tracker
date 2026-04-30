package service

import (
	"context"
	"money-tracker/api/internal/user"
)

type CreateUser struct {
	repo user.UserRepository
}

func (c *CreateUser) Execute(ctx context.Context, user user.User) (string, error) {
	return c.repo.Save(ctx, user)
}
