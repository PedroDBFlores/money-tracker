package service_test

import (
	"context"
	"money-tracker/api/internal/user"
	"money-tracker/api/internal/user/service"

	"money-tracker/api/internal/user/mocks"
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	t.Run("creates with success", func(t *testing.T) {
		user := user.User{
			Name:         "John Doe",
			Email:        "john@example.com",
			PasswordHash: "the_best_hash",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		repo := &mocks.MockUserRepository{}
		repo.On("Save", mock.Anything, user).Return("12", nil)

		service := service.NewUserService(repo)
		id, err := service.CreateUser.Execute(context.TODO(), user)
		assert.NoError(t, err)
		assert.Equal(t, "12", id)

		repo.AssertCalled(t, "Save", mock.Anything, user)
	})

	t.Run("returns error when save fails", func(t *testing.T) {
		repo := &mocks.MockUserRepository{}
		repo.On("Save", mock.Anything, mock.Anything).Return("", assert.AnError)

		service := service.NewUserService(repo)
		_, err := service.CreateUser.Execute(context.TODO(), user.User{})
		assert.Error(t, err)
	})
}
