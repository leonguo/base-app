package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"echo-app/api/controllers"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	app.OnErrorCode(iris.StatusNotFound,notFoundHandler)

	app.Controller("/user",new(controllers.UserController))

	app.Get("ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})
	authMiddleware := func(ctx iris.Context) {
		ctx.Next()
	}
	v1 := app.Party("/v1", authMiddleware)
	usersAPI := v1.Party("/users")
	{
		// http://localhost:8080/api/users
		usersAPI.Get("/", h)
		usersAPI.Post("/", h)
		// http://localhost:8080/api/users/42
		usersAPI.Get("/{userid:int}", func(ctx iris.Context) {
			ctx.Writef("user with id: %s", ctx.Params().Get("userid"))
		})
	}
	app.Run(iris.Addr("localhost:8080"))
}

func notFoundHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{"message":"not found"})
}

func h(ctx iris.Context) {
	method := ctx.Method()       // the http method requested a server's resource.
	subdomain := ctx.Subdomain() // the subdomain, if any.

	// the request path (without scheme and host).
	path := ctx.Path()
	// how to get all parameters, if we don't know
	// the names:
	paramsLen := ctx.Params().Len()

	ctx.Params().Visit(func(name string, value string) {
		ctx.Writef("%s = %s\n", name, value)
	})
	ctx.Writef("Info\n\n")
	ctx.Writef("Method: %s\nSubdomain: %s\nPath: %s\nParameters length: %d", method, subdomain, path, paramsLen)
}