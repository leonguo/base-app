package controllers

import (
	//"echo-app/models"
	"base-app/services"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"base-app/models"
)

// VideoController is our /user controller.
// VideoController is responsible to handle the following requests:
// GET  			/user/register
// POST 			/user/register
// GET 				/user/login
// POST 			/user/login
// GET 				/user/me
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

// PostRegister handles POST: http://localhost:8080/user/register.
func (c *VideoController) PostRegister() mvc.Result {
	// get firstname, username and password from the form.
	//var (
	//	username  = c.Ctx.FormValue("username")
	//	password  = c.Ctx.FormValue("password")
	//)

	// create the new user, the password will be hashed by the service.
	//u, err := c.Service.Create(password, models.User{
	//	Username:  username
	//})

	// set the user's id to this session even if err != nil,
	// the zero id doesn't matters because .getCurrentUserID() checks for that.
	// If err != nil then it will be shown, see below on mvc.Response.Err: err.

	return mvc.Response{
		// if not nil then this error will be shown instead.
		//Err: err,
		// redirect to /user/me.
		Path: "/user/me",
		// When redirecting from POST to GET request you -should- use this HTTP status code,
		// however there're some (complicated) alternatives if you
		// search online or even the HTTP RFC.
		// Status "See Other" RFC 7231, however iris can automatically fix that
		// but it's good to know you can set a custom code;
		// Code: 303,
	}

}

// GetLogin handles GET: http://localhost:8080/user/login.
func (c *VideoController) GetLogin() mvc.Result {
	if c.isLoggedIn() {
		// if it's already logged in then destroy the previous session.
	}
	c.Ctx.Application().Logger().Warnf("get user info ")
	return mvc.Response{Code: 200, Object: models.User{ID: 2, Username: "DDD"}}
}

func (c *VideoController) GetMe() mvc.Result {
	if c.isLoggedIn() {
		// if it's already logged in then destroy the previous session.
	}
	return mvc.Response{Code: 200, Object: models.User{ID: 2, Username: "222"}}
}
