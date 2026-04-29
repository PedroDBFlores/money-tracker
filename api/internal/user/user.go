package user

import "time"

type User struct {
	Id           int64
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUser(id int64, name, email, passwordHash string, createdAt, updatedAt time.Time) User {
	return User{
		Id:           id,
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
