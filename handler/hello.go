package handler

// mỗi một api sẽ gọi đến 1 function handler
// đây là package chứa các function handler
import (
	"myapp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	// the way to use struct of models package
	x := &models.X{
		Text: "Hello world!",
	}
	return c.JSON(http.StatusOK, x)
}
