package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"fmt"
)

func main() {
	server := echo.New()

	server.GET("/", hello)

	server.POST("/login", login)

	server.Logger.Fatal(server.Start(":8888"))

}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world!")
}

func login(c echo.Context) error {
	req := new(LoginRequest)
	c.Bind(req)

	fmt.Printf("req:%v", req)

	if (req.Username != "admin" || req.Password != "123") {
		return c.String(http.StatusUnauthorized, "username or password is incorrect")
	}

	return c.JSON(http.StatusOK, &LoginResponse{Token: "123456"})
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}