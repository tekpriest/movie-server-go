package utils

import "github.com/en1tan/movie-server/models"

type SuccessResponse struct {
	Status     bool   `json:"status" default:"true"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
} // @Name SuccessResponse

type MetaData struct {
	Total int `json:"total"`
	Page  int `json:"page"`
}

type SuccessReponseForArrayOfMovies struct {
	SuccessResponse
	Data *[]models.Movie
	MetaData
} // @Name SuccessReponseForArrayOfMovies

type SuccesResponseForMovie struct {
	SuccessResponse
	Data *models.Movie
} // @Name SuccessResponseForMovie
