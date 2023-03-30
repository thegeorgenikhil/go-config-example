package config

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type AppConfig struct {
	DatabaseURI  string
	DatabaseName string
	Client       *mongo.Client
}
