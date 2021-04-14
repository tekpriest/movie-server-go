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
	List(page, limit int) (*[]models.Movie, int64, error)
	GetByID(id string) (*models.Movie, error)
	Create(m models.Movie) (*models.Movie, error)
	UpdateByID(id string, m models.Movie) (*models.Movie, error)
	DeleteByID(id string) error
}

type movieService struct {
	db *mongo.Collection
}

func NewMovieService(conn database.DatabaseConnection) MovieService {
	return &movieService{db: conn.Get().Collection("movies")}
}

func (ms *movieService) List(page, limit int) (*[]models.Movie, int64, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var movies []models.Movie
	skip := int64((page - 1) * limit)
	opts := options.Find()
	opts.SetSort(bson.D{{Key: "createdAt", Value: -1}})
	opts.SetLimit(int64(limit) * 1)
	opts.SetSkip(skip)
	cur, _ := ms.db.Find(ctx, bson.D{{}}, opts)
	count, err := ms.db.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return nil, 0, err
	}
	if err = cur.All(ctx, &movies); err != nil {
		return nil, 0, err
	}
	return &movies, count, err
}

func (ms *movieService) Create(m models.Movie) (*models.Movie, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	m.CreatedAt = time.Now()
	newMovie, err := ms.db.InsertOne(ctx, &m)
	if err != nil {
		return nil, err
	}
	if err := ms.db.FindOne(ctx, bson.M{"_id": newMovie.InsertedID}).Decode(&m); err != nil {
		return nil, err
	}
	return &m, err
}

func (ms *movieService) GetByID(id string) (*models.Movie, error) {
	var movie models.Movie
	movieID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := ms.db.FindOne(ctx, bson.M{"_id": movieID}).Decode(&movie); err != nil {
		return nil, err
	}
	return &movie, err
}

func (ms *movieService) UpdateByID(id string, m models.Movie) (*models.Movie, error) {
	movieID, e := primitive.ObjectIDFromHex(id)
	if e != nil {
		return nil, e
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	m.UpdatedAt = time.Now()
	err := ms.db.FindOneAndReplace(ctx, bson.M{"_id": movieID}, &m).Decode(&m)
	if err != nil {
		return nil, err
	}
	return &m, err
}

func (ms *movieService) DeleteByID(id string) error {
	movieID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, deleteErr := ms.db.DeleteOne(ctx, bson.M{"_id": movieID})
	if err != nil {
		return deleteErr
	}
	return err
}
