package controllers

import (
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/en1tan/movie-server/models"
	"github.com/en1tan/movie-server/services"
	"github.com/en1tan/movie-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// controllers Movie Controller

type MovieController interface {
	ListMovies(c *gin.Context)
	GetMovieByID(c *gin.Context)
	CreateMovie(c *gin.Context)
	UpdateMovieByID(c *gin.Context)
	DeleteMovie(c *gin.Context)
}

type movieController struct {
	ms services.MovieService
	vu utils.ValidatorUtil
}

func NewMovieController(ms services.MovieService, vu utils.ValidatorUtil) MovieController {
	return &movieController{ms, vu}
}

// @ListMovies godoc
// @Summary Lists all the movies
// @Param page query string false "page"
// @Param limit query string false "limit"
// @Success 200 {array} utils.SuccessReponseForArrayOfMovies
// @Failure 400 {object} utils.ApiError
// @Router /movies [get]
// @Tags Movie
// @ID ListMovies
func (mc *movieController) ListMovies(c *gin.Context) {
	page, limit := 1, 20
	pageQuery := c.Query("page")
	limitQuery := c.Query("limit")
	if pageQuery != "" || limitQuery != "" {
		p, err := strconv.Atoi(pageQuery)
		l, err := strconv.Atoi(limitQuery)
		if err != nil {
			c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("invalid page query parameter")))
			return
		}
		page = p
		limit = l
	}
	movies, count, err := mc.ms.List(page, limit)
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  true,
			"message": "Fetched All movies",
			"data": gin.H{
				"movies": &movies,
				"page":   page,
				"total":  math.Ceil(float64(count) / float64(limit)),
			},
		},
	)
}

// @CreateMovie godoc
// @Param Movie body models.Movie true "New Movie"
// @Summary Create New Movie
// @Success 201 {object} utils.SuccesResponseForMovie
// @Failure 400 {object} utils.ApiError
// @Failure 500 {object} utils.ApiError
// @Tags Movie
// @Router /movies [post]
// @ID CreateMovie
func (mc *movieController) CreateMovie(c *gin.Context) {
	var m models.Movie
	if err := c.ShouldBindJSON(&m); err != nil {
		var v validator.ValidationErrors
		if errors.As(err, &v) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": mc.vu.CreateMovieValidator(v)})
			return
		}
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

// @GetMovieByID godoc
// @Param id path string true "Movie ID"
// @Summary Get Movie by ID
// @Success 200 {object} utils.SuccesResponseForMovie
// @Failure 404 {object} utils.ApiError
// @Failure 400 {object} utils.ApiError
// @Failure 500 {object} utils.ApiError
// @Tags Movie
// @Router /movies/{id} [get]
// @ID GetMovieByID
func (mc *movieController) GetMovieByID(c *gin.Context) {
	movie, err := mc.ms.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(utils.ErrorFromDatabase(err))
	}
	c.JSON(http.StatusOK, movie)
}

// @UpdateMovieByID godoc
// @Param id path string true "Movie ID"
// @Param Movie body models.Movie true "Movie"
// @Summary Update Movie Details
// @Success 200 {object} utils.SuccesResponseForMovie
// @Failure 404 {object} utils.ApiError
// @Failure 400 {object} utils.ApiError
// @Failure 500 {object} utils.ApiError
// @Tags Movie
// @Router /movies/{id} [put]
// @ID UpdateMovieByID
func (mc *movieController) UpdateMovieByID(c *gin.Context) {
	var m models.Movie
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("wrong request body")))
		return
	}
	updatedMovie, err := mc.ms.UpdateByID(c.Param("id"), m)
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("an error occured")))
	}
	c.JSON(http.StatusOK, updatedMovie)
}

// @DeleteMovie godoc
// @Param id path string true "Movie ID"
// @Summary Delete Movie by ID
// @Success 200 {string} string "Movie Deleted"
// @Failure 404 {object} utils.ApiError
// @Failure 400 {object} utils.ApiError
// @Failure 500 {object} utils.ApiError
// @Tags Movie
// @Router /movies/{id} [delete]
// @ID DeleteMovie
func (mc *movieController) DeleteMovie(c *gin.Context) {
	err := mc.ms.DeleteByID(c.Param("id"))
	if err != nil {
		c.JSON(utils.CreateApiError(http.StatusBadRequest, errors.New("an error occured")))
		return
	}
	c.JSON(http.StatusOK, "Movie Deleted")
}
