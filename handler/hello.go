package handler

// mỗi một api sẽ gọi đến 1 function handler
// đây là package chứa các function handler
import (
	"fmt"
	"myapp/models"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	//https://echo.labstack.com/cookbook/jwt/#server-using-a-user-defined-keyfunc
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	username := claims["username"].(string)
	admin := claims["admin"].(bool)

	message := fmt.Sprintf("Hello %s, Are you admin? %v", username, admin)

	// the way to use struct of models package
	x := &models.X{
		Text: message,
	}
	return c.JSON(http.StatusOK, x)
}
