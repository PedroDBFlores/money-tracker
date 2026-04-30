package service_test

import (
	"context"
	"money-tracker/api/internal/user"
	"money-tracker/api/internal/user/mocks"
	"money-tracker/api/internal/user/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetByEmail(t *testing.T) {
	t.Run("should return user by email if found", func(t *testing.T) {
		expectedUser := user.User{
			Id:           123,
			Name:         "John Doe",
			Email:        "john@example.com",
			PasswordHash: "a_nice_hash",
			CreatedAt:    time.Time{},
			UpdatedAt:    time.Time{},
		}
		repo := mocks.MockUserRepository{}
		repo.On("GetByEmail", mock.Anything, "john@example.com").Return(&expectedUser, true, nil)
		service := service.NewUserService(&repo)

		actualUser, found, err := service.GetByEmail.Execute(context.TODO(), "john@example.com")
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, *actualUser)
		assert.True(t, found)

		repo.AssertCalled(t, "GetByEmail", mock.Anything, "john@example.com")
	})

	t.Run("should return nil and false if user not found", func(t *testing.T) {
		repo := mocks.MockUserRepository{}
		repo.On("GetByEmail", mock.Anything, mock.Anything).Return(nil, false, nil)
		service := service.NewUserService(&repo)

		actualUser, found, err := service.GetByEmail.Execute(context.TODO(), "john@example.com")
		assert.NoError(t, err)
		assert.Nil(t, actualUser)
		assert.False(t, found)

		repo.AssertCalled(t, "GetByEmail", mock.Anything, "john@example.com")
	})

	t.Run("should return an error if getting the user fails", func(t *testing.T) {
		repo := mocks.MockUserRepository{}
		repo.On("GetByEmail", mock.Anything, mock.Anything).Return(nil, false, assert.AnError)
		service := service.NewUserService(&repo)

		actualUser, found, err := service.GetByEmail.Execute(context.TODO(), "john@example.com")
		assert.Error(t, err)
		assert.Nil(t, actualUser)
		assert.False(t, found)

		repo.AssertCalled(t, "GetByEmail", mock.Anything, "john@example.com")
	})
}
