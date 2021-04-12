/* Package controllers
@package controllers
*/

package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/en1tan/movie-server/models"
	"github.com/en1tan/movie-server/services"
	"github.com/en1tan/movie-server/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Package Controllers

type MovieController interface {
	ListMovies(c *gin.Context)
	GetMovieByID(c *gin.Context)
	CreateMovie(c *gin.Context)
	UpdateMovieByID(c *gin.Context)
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
	pageQuery := c.Query("page")
	if pageQuery != "" {
		p, err := strconv.Atoi(pageQuery)
		if err != nil {
			c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("invalid page query parameter")))
			return
		}
		page = p
	}
	movies, err := mc.ms.List(page)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}
	c.JSON(http.StatusOK, movies)
}

func (mc *movieController) CreateMovie(c *gin.Context) {
	var m models.Movie
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("wrong request body")))
		return
	}
	movie, err := mc.ms.Create(m)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}
	c.JSON(http.StatusCreated, movie)
}

func (mc *movieController) GetMovieByID(c *gin.Context) {
	movie, err := mc.ms.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
	}
	c.JSON(http.StatusOK, movie)
}

func (mc *movieController) UpdateMovieByID(c *gin.Context) {
	var m models.Movie
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("wrong request body")))
		return
	}
	updatedMovie, err := mc.ms.UpdateByID(c.Param("id"), m)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, updatedMovie)
}

func (mc *movieController) DeleteMovie(c *gin.Context) {
	err := mc.ms.DeleteByID(c.Param("id"))
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("an error occured")))
		return
	}
	c.JSON(http.StatusOK, "Movie Deleted")
}
