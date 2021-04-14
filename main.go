package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/en1tan/movie-server/config"
	"github.com/en1tan/movie-server/controllers"
	"github.com/en1tan/movie-server/database"
	_ "github.com/en1tan/movie-server/docs"
	"github.com/en1tan/movie-server/routes"
	"github.com/en1tan/movie-server/services"
	"github.com/en1tan/movie-server/utils"
	_ "github.com/en1tan/movie-server/utils"
)

// @title Movie Server API
// @version 1.0
// @description Movie Server written in go
// @termsOfService http://#

// @license.name MIT

// @Accept json
// @Produce json
// @basePath /api

// @Failure 404 {object} utils.ApiError
// @Failure 400 {object} utils.ApiError
// @Failure 500 {object} utils.ApiError

func main() {
	c := config.NewConfig()

	r := routes.NewRouter(c)
	conn := database.NewDatabaseConnection(c)

	ms := services.NewMovieService(conn)
	vu := utils.NewValidatorUtil()
	mc := controllers.NewMovieController(ms, *vu)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.RegisterMovieRoutes(mc)
	r.Serve()
}
