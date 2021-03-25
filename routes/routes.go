package routes

import "github.com/en1tan/movie-server/controllers"

func (r *router) RegisterMovieRoutes(c controllers.MovieController) {
	rg := r.Group("api/movies")
	rg.GET("/", c.ListMovies)
	rg.POST("/", c.CreateMovie)
	rg.GET("/:id", c.GetMovieById)
	rg.PUT("/:id", c.UpdateMovieById)
	rg.DELETE("/:id", c.DeleteMovie)
}
