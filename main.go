package main

import (
	// the way to import package in this project

	"myapp/handler"
	"myapp/mdw"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	server := echo.New()

	// middleware: key word import by "github.com/labstack/echo/v4/middleware"
	// this middleware uses for all API Call
	// log request api history
	server.Use(middleware.Logger())

	// how to use middleware for a api method??
	// api request -> middleware run first
	// pass-> run this login controller
	// this field is check in Basic Auth (Auth tab in post man)
	// this solution use the middleware only for this api
	server.POST("/login", handler.Login, middleware.BasicAuth(mdw.BasicAuth))

	// check jwt is valid
	isLogin := middleware.JWT([]byte("secret_key"))

	IsAdminMdw := mdw.IsAdminMdw

	// it make a check " "missing or malformed jwt""
	server.GET("/", handler.Hello, isLogin)

	// the middleware function list: execute in order
	server.GET("/admin", handler.Hello, isLogin, IsAdminMdw)

	// create group api
	groupv2 := server.Group("/v2")
	groupv2.GET("/hello2", handler.Hello2)
	groupv2.GET("/hello3", handler.Hello3)

	// use middleware function for this whole api group
	groupUser := server.Group("/api/user")
	groupUser.GET("/user1", handler.User)
	groupUser.GET("/user2", handler.User2)
	groupUser.GET("/get-all", handler.GetAllUsers)

	server.Logger.Fatal(server.Start(":8888"))

}
