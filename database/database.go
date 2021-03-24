package database

import (
	"context"
	"time"

	"github.com/en1tan/movie-server/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConnection interface {
	Get() *mongo.Database
}

type databaseConnection struct {
	DB *mongo.Database
}

func NewDatabaseConnection(c *config.Config) DatabaseConnection {
	config := c.Get()
	databaseUri := config.GetString("db.uri")

	client, err := mongo.NewClient(options.Client().ApplyURI(databaseUri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("movie_server")
	if err != nil {
		log.Fatal(err)
	}

	return &databaseConnection{DB: db}
}

func (d *databaseConnection) Get() *mongo.Database {
	return d.DB
}
