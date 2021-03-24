package controllers

import (
	"net/http"

	"github.com/en1tan/movie-server/models"
	"github.com/en1tan/movie-server/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type MovieController interface {
	ListMovies(c *gin.Context)
	CreateMovie(c *gin.Context)
}

type movieController struct {
	ms services.MovieService
}

func NewMovieController(ms services.MovieService) MovieController {
	return &movieController{ms}
}

func (mc *movieController) ListMovies(c *gin.Context) {
	page := 0
	err, movies := mc.ms.List(page)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, movies)
	return
}

func (mc *movieController) CreateMovie(c *gin.Context) {
	var m models.Movie
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(400, "An error occured")
		return
	}
	err, movie := mc.ms.Create(m)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusCreated, movie)
	return
}
