package mocks

import (
	"context"
	"money-tracker/api/internal/user"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(ctx context.Context, user user.User) (string, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(string), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*user.User, bool, error) {
	args := m.Called(ctx, email)
	var usr *user.User
	if args.Get(0) != nil {
		usr = args.Get(0).(*user.User)
	}

	return usr, args.Get(1).(bool), args.Error(2)
}
