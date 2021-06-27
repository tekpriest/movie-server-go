package utils

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApiError struct {
	Status     bool   `json:"status" default:"false"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
} // @Name ResponseError

func CreateApiError(statusCode int, err error) (int, *ApiError) {
	log.Error(err.Error())
	status := false
	message := err.Error()
	return statusCode, &ApiError{
		Status:     status,
		StatusCode: statusCode,
		Message:    message,
	}
}

func ErrorFromDatabase(err error) (int, *ApiError) {
	switch err {
	case mongo.ErrNoDocuments:
		return CreateApiError(http.StatusNotFound, err)
	default:
		return CreateApiError(http.StatusInternalServerError, err)
	}
}
