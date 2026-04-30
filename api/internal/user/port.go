package user

import "context"

// UserRepository defines the interface for user repository operations
type UserRepository interface {
	Save(ctx context.Context, u User) (string, error)
	GetByEmail(ctx context.Context, email string) (*User, bool, error)
}
