package handler

import (
	"money-tracker/api/internal/user/service"

	"github.com/gofiber/fiber/v3"
)

type CreateUser struct {
	service service.CreateUser
}

func NewCreateUser(service service.CreateUser) *CreateUser {
	return &CreateUser{service: service}
}

func (h *CreateUser) Handle(ctx fiber.Ctx) error {
	ctx.Status(fiber.StatusCreated)

	return nil
}
