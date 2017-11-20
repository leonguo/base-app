package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"echo-app/api/controllers"
	_ "github.com/lib/pq"
	"github.com/go-xorm/xorm"
	"time"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)

	app.Get("ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})
	authMiddleware := func(ctx iris.Context) {
		ctx.Next()
	}
	v1 := app.Party("/v1", authMiddleware)
	usersAPI := v1.Party("/user")
	{
		usersAPI.Controller("/", new(controllers.UserController))
	}
	// 节目
	videoAPI := v1.Party("/video")
	{
		videoAPI.Get("/test", h)
		videoAPI.Controller("/", new(controllers.VideoController))
	}

	app.Get("/db/get", func(ctx iris.Context) {
		conn := "postgres://postgres:123456@120.24.229.18/app?sslmode=disable"
		orm, err := xorm.NewEngine("postgres", conn)
		if err != nil {
			app.Logger().Fatalf("orm failed to initialized: %v", err)
		}

		iris.RegisterOnInterrupt(func() {
			orm.Close()
		})

		type Users struct {
			Uid       int64
			Salt      string
			Username  string
			Password  string    `xorm:"varchar(200)"`
			CreatedAt time.Time `xorm:"created_at"`
			UpdatedAt time.Time `xorm:"updated_at"`
		}
		err = orm.Sync2(new(Users))
		user := Users{Uid: 2}
		if ok, _ := orm.Get(&user); ok {
			ctx.Writef("user found: %v", user)
		}
		app.Logger().Info("user found username : ", user.Username)

	})
	app.Run(iris.Addr("localhost:8080"))
}

func notFoundHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "not found"})
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
