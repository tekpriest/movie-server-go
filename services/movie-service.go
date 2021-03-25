package services

import (
	"context"
	"log"
	"time"

	"github.com/en1tan/movie-server/database"
	"github.com/en1tan/movie-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MovieService interface {
	List(page int) (error, *[]models.Movie)
	GetById(id string) (error, *models.Movie)
	Create(m models.Movie) (error, *models.Movie)
	UpdateById(id string, m models.Movie) (error, *models.Movie)
	DeleteById(id string) error
}

type movieService struct {
	db *mongo.Collection
}

func NewMovieService(conn database.DatabaseConnection) MovieService {
	return &movieService{db: conn.Get().Collection("movies")}
}

func (ms *movieService) List(page int) (error, *[]models.Movie) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var movies []models.Movie
	opts := options.Find()
	opts.SetSort(bson.D{{"createdAt", -1}})
	cur, err := ms.db.Find(ctx, bson.D{}, opts)
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

func (ms *movieService) GetById(id string) (error, *models.Movie) {
	var movie models.Movie
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := ms.db.FindOne(ctx, bson.M{}).Decode(&movie); err != nil {
		log.Fatal(err)
	}
	return nil, &movie
}

func (ms *movieService) UpdateById(id string, m models.Movie) (error, *models.Movie) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := ms.db.UpdateByID(ctx, id, bson.M{"$set": m})
	if err != nil {
		log.Fatal(err)
	}
	return err, &m
}

func (ms *movieService) DeleteById(id string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := ms.db.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
	return err
}
