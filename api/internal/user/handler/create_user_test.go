package handler_test

import (
	"money-tracker/api/internal/user/handler"
	"money-tracker/api/internal/user/mocks"
	"money-tracker/api/internal/user/service"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	t.Run("returns 201 when creating the user successfully", func(t *testing.T) {
		repo := mocks.MockUserRepository{}
		service := service.NewUserService(&repo)
		handler := handler.NewCreateUser(service.CreateUser)
		app := fiber.New()
		app.Post("/users", handler.Handle)
		req, err := http.NewRequest("POST", "/users", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Body = http.NoBody
		req.Header.Set("Content-Type", "application/json")
		res, _ := app.Test(req)

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
