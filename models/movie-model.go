package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Title     string             `bson:"title,omitempty" json:"title"`
	Poster    string             `bson:"poster" json:"poster"`
	Link      string             `bson:"link,omitempty" json:"link"`
	Type      string             `bson:"type" json:"type"`
	Year      string             `bson:"year,omitempty" json:"year"`
	ImdbId    string             `bson:"imdbid,omitempty" json:"imdbid"`
	Genre     string             `bson:"genre" json:"genre"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt" bson:default:now()`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
