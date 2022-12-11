package handler

// mỗi một api sẽ gọi đến 1 function handler
// đây là package chứa các function handler
import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func User2(c echo.Context) error {

	return c.String(http.StatusOK, "user 2")
}

func User(c echo.Context) error {

	return c.String(http.StatusOK, "user 1")
}
