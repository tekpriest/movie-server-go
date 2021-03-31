package routes

import "github.com/en1tan/movie-server/controllers"

func (r *router) RegisterMovieRoutes(c controllers.MovieController) {
	rg := r.Group("api/movies")
	rg.GET("/", c.ListMovies)
	rg.POST("/", c.CreateMovie)
	rg.GET("/:id", c.GetMovieByID)
	rg.PUT("/:id", c.UpdateMovieByID)
	rg.DELETE("/:id", c.DeleteMovie)
}
