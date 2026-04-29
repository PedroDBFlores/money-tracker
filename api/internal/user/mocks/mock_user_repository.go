package mocks

import (
	"money-tracker/api/internal/user"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(user user.User) (int64, error) {
	args := m.Called(user)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(email string) (*user.User, bool, error) {
	args := m.Called(email)
	return args.Get(0).(*user.User), args.Get(1).(bool), args.Error(2)
}
