package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id",
		func(c echo.Context) error {
			u := new(User)
			id := c.Param("id")
			newId,err:=strconv.Atoi(id)
			if err != nil {
				c.Logger().Debug(err)
			}
			u.Id = newId
			return c.JSON(http.StatusOK, u)
		})
	//e.PUT("/users/:id", updateUser)
	e.Logger.Fatal(e.Start(":3000"))
}
