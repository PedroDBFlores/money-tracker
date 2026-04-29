package database

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Database struct {
	Client *mongo.Client
}

func NewDatabase(config Config) (*Database, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(config.URI))
	if err != nil {
		return nil, err
	}
	return &Database{Client: client}, nil
}

func (db *Database) Close() error {
	return db.Close()
}
