package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/basicauth"
	"github.com/kataras/iris/mvc"
)

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

var basicAuth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"admin": "password",
	},
})

func main() {
	app := iris.New()
	app.Use(basicauth)
	app.Controller("/movies", new(MovieController))
	app.Run(iris.Addr(":8080"))
}

type MovieController struct {
	mvc.C //继承
}

func (c *MoviesController) Get() []Movie {
	return movies
}

func (c *MoviesController) GetBy(id int) Movie {
	return movies[id]
}

func (c *MoviesController) PutBy(id int) Movie {
	m := movies[id]
}

//...
