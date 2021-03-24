package services

import (
	"context"
	"time"

	"github.com/en1tan/movie-server/database"
	"github.com/en1tan/movie-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieService interface {
	List(page int) (error, *[]models.Movie)
	// GetById(id string) (error, *models.Movie)
	Create(m models.Movie) (error, *models.Movie)
}

type movieService struct {
	db *mongo.Collection
}

func NewMovieService(conn database.DatabaseConnection) MovieService {
	return &movieService{db: conn.Get().Collection("movies")}
}

func (ms *movieService) List(page int) (error, *[]models.Movie) {
	var movies []models.Movie
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := ms.db.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	if err = cur.All(ctx, &movies); err != nil {
		panic(err)
	}
	return err, &movies
}

func (ms *movieService) Create(m models.Movie) (error, *models.Movie) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := ms.db.InsertOne(ctx, &m)
	if err != nil {
		panic(err)
	}
	return err, &m
}
