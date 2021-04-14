package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Movie Model
type Movie struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id" readonly:"true"`
	Title     string             `bson:"title,omitempty" binding:"required" json:"title"`
	Poster    string             `bson:"poster" binding:"required" json:"poster"`
	Link      string             `bson:"link" binding:"required" json:"link"`
	Type      string             `bson:"type" binding:"required" json:"type"`
	Year      string             `bson:"year" binding:"required" json:"year"`
	ImdbID    string             `bson:"imdbid" binding:"required" json:"imdbid"`
	Genre     string             `bson:"genre" binding:"required" json:"genre"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt" readonly:"true"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty" readonly:"true"`
} // @Name Movie
