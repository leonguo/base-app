package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// Movie is our sample data structure.
type Movie struct {
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
	Poster string `json:"poster"`
}

var movies = []Movie{
	{
		Name:   "Casablanca",
		Year:   1942,
		Genre:  "Romance",
		Poster: "https://iris-go.com/images/examples/mvc-movies/1.jpg",
	},
	{
		Name:   "Gone with the Wind",
		Year:   1939,
		Genre:  "Romance",
		Poster: "https://iris-go.com/images/examples/mvc-movies/2.jpg",
	},
	{
		Name:   "Citizen Kane",
		Year:   1941,
		Genre:  "Mystery",
		Poster: "https://iris-go.com/images/examples/mvc-movies/3.jpg",
	},
	{
		Name:   "The Wizard of Oz",
		Year:   1939,
		Genre:  "Fantasy",
		Poster: "https://iris-go.com/images/examples/mvc-movies/4.jpg",
	},
}

func main() {
	app := iris.New()
	app.Controller("/hello",new(MoviesController))
	app.Run(iris.Addr("localhost:8080"))
}

// MoviesController is our /movies controller.
type MoviesController struct {
	mvc.C
}

// Get returns list of the movies
// Demo:
// curl -i http://localhost:8080/movies
func (c *MoviesController) Get() []Movie {
	return movies
}