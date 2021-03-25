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
	GetMovieById(c *gin.Context)
	CreateMovie(c *gin.Context)
	UpdateMovieById(c *gin.Context)
	DeleteMovie(c *gin.Context)
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

func (mc *movieController) GetMovieById(c *gin.Context) {
	err, movie := mc.ms.GetById(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, movie)
	return
}

func (mc *movieController) UpdateMovieById(c *gin.Context) {
	var m models.Movie
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(400, "An Error Occured")
		return
	}
	err, updatedMovie := mc.ms.UpdateById(c.Param("id"), m)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, updatedMovie)
}

func (mc *movieController) DeleteMovie(c *gin.Context) {
	err := mc.ms.DeleteById(c.Param("id"))
	if err != nil {
		c.JSON(400, "An Error Occured")
		return
	}
	c.JSON(http.StatusOK, "Movie Deleted")
}
