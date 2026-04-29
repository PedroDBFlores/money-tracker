package adapters

import (
	"money-tracker/api/internal/user"
	"money-tracker/api/pkg/database"
)

type UserRepository struct {
	db *database.Database
}

func NewUserRepository(db *database.Database) UserRepository {
	return UserRepository{db: db}
}

func (r *UserRepository) Save(u *user.User) (int64, error) {
	panic("implement me")
}

func (r *UserRepository) GetByEmail(email string) (*user.User, bool, error) {
	panic("implement me")
}
