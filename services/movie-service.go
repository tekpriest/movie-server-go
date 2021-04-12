package services

import (
	"context"
	"time"

	"github.com/en1tan/movie-server/database"
	"github.com/en1tan/movie-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	opts.SetLimit(40)
	cur, err := ms.db.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		return err, nil
	}
	if err = cur.All(ctx, &movies); err != nil {
		return err, nil
	}
	return err, &movies
}

func (ms *movieService) Create(m models.Movie) (error, *models.Movie) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	m.CreatedAt = time.Now()
	newMovie, err := ms.db.InsertOne(ctx, &m)
	if err != nil {
		return err, nil
	}
	if err := ms.db.FindOne(ctx, bson.M{"_id": newMovie.InsertedID}).Decode(&m); err != nil {
		return err, nil
	}
	return err, &m
}

func (ms *movieService) GetById(id string) (error, *models.Movie) {
	var movie models.Movie
	movieId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err, nil
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := ms.db.FindOne(ctx, bson.M{"_id": movieId}).Decode(&movie); err != nil {
		return err, nil
	}
	return nil, &movie
}

func (ms *movieService) UpdateById(id string, m models.Movie) (error, *models.Movie) {
	movieID, e := primitive.ObjectIDFromHex(id)
	if e != nil {
		return e, nil
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	m.UpdatedAt = time.Now()
	err := ms.db.FindOneAndReplace(ctx, bson.M{"_id": movieID}, &m).Decode(&m)
	if err != nil {
		return err, nil
	}
	return err, &m
}

func (ms *movieService) DeleteById(id string) error {
	movieId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, deleteErr := ms.db.DeleteOne(ctx, bson.M{"_id": movieId})
	if err != nil {
		return deleteErr
	}
	return err
}
