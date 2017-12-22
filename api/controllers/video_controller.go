package controllers

import (
	//"echo-app/models"
	"base-app/services"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"base-app/models"
)

// VideoController is our /video controller.
// All HTTP Methods /user/logout
type VideoController struct {
	// mvc.C is just a lightweight lightweight alternative
	// to the "mvc.Controller" controller type,
	// use it when you don't need mvc.Controller's fields
	// (you don't need those fields when you return values from the method functions).
	mvc.C

	// Our UserService, it's an interface which
	// is binded from the main application.
	Service services.UserService
}

// BeginRequest will set the current session to the controller.
//
// Remember: iris.Context and context.Context is exactly the same thing,
// iris.Context is just a type alias for go 1.9 users.
// We use context.Context here because we don't need all iris' root functions,
// when we see the import paths, we make it visible to ourselves that this file is using only the context.
func (c *VideoController) BeginRequest(ctx context.Context) {
	c.C.BeginRequest(ctx)
}

func (c *VideoController) getCurrentUserID() int64 {
	return 1
}

func (c *VideoController) isLoggedIn() bool {
	return c.getCurrentUserID() > 0
}

func (c *VideoController) GetVideoInfo() mvc.Result {
	if c.isLoggedIn() {
		// if it's already logged in then destroy the previous session.
	}
	return mvc.Response{Code: 200, Object: models.User{ID: 2, Username: "222"}}
}
