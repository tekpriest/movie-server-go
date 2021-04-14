package utils

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
} // @Name ResponseError

func CreateApiError(status int, err error) (int, *ApiError) {
	log.Error(err.Error())
	message := err.Error()
	return status, &ApiError{
		Status:  status,
		Message: message,
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
