package adapters

import (
	"context"
	"money-tracker/api/internal/user"
	"money-tracker/api/pkg/database"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type mongoDBUser struct {
	ID           bson.ObjectID `bson:"_id"`
	Name         string        `bson:"name"`
	Email        string        `bson:"email"`
	PasswordHash string        `bson:"passwordHash"`
	CreatedAt    time.Time     `bson:"createdAt"`
	UpdatedAt    time.Time     `bson:"updatedAt"`
}

func toMongoDBUser(u *user.User) mongoDBUser {
	return mongoDBUser{
		ID:           bson.NewObjectID(),
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

type UserRepository struct {
	db *database.Database
}

func NewUserRepository(db *database.Database) UserRepository {
	return UserRepository{db: db}
}

func (r *UserRepository) Save(ctx context.Context, u *user.User) (string, error) {
	db := r.db.Client.Database("money-tracker")
	dbUser := toMongoDBUser(u)
	_, err := db.Collection("users").InsertOne(ctx, dbUser)
	if err != nil {
		return "", err
	}

	return dbUser.ID.String(), nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*user.User, bool, error) {
	panic("implement me")
}
