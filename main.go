package main

import (
	"github.com/en1tan/movie-server/config"
	"github.com/en1tan/movie-server/controllers"
	"github.com/en1tan/movie-server/database"
	"github.com/en1tan/movie-server/routes"
	"github.com/en1tan/movie-server/services"
)

func main() {
	c := config.NewConfig()

	r := routes.NewRouter(c)
	conn := database.NewDatabaseConnection(c)

	ms := services.NewMovieService(conn)
	mc := controllers.NewMovieController(ms)

	r.RegisterMovieRoutes(mc)
	r.Serve()
}
