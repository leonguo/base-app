package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id",
		func(c echo.Context) error {
			u := new(User)
			id := c.Param("id")
			name := c.QueryParam("name")
			c.Logger().Error(name)
			newId, err := strconv.Atoi(id)
			if err != nil {
				c.Logger().Debug(err)
			}
			u.Id = newId
			u.Name = name
			return c.JSON(http.StatusOK, u)
		})
	e.GET("/error", func(context echo.Context) error {
		return echo.NewHTTPError(http.StatusBadRequest)
	})
	//e.PUT("/users/:id", updateUser)

	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))
	e.Logger.Fatal(e.Start(":3000"))
}
