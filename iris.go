package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"echo-app/api/controllers"
)

// Movie is our sample data structure.
type Movie struct {
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
	Poster string `json:"poster"`
}

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	app.OnErrorCode(iris.StatusNotFound,notFoundHandler)

	app.Controller("/user",new(controllers.UserController))

	app.Get("ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("test", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message":"hello iris"})
	})
	app.Run(iris.Addr("localhost:8080"))
}

func notFoundHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{"message":"not found"})
}