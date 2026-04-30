package service

import (
	"context"
	"money-tracker/api/internal/user"
)

type GetByEmail struct {
	repo user.UserRepository
}

func (g GetByEmail) Execute(ctx context.Context, email string) (*user.User, bool, error) {
	return g.repo.GetByEmail(ctx, email)
}
