package handler

// mỗi một api sẽ gọi đến 1 function handler
// đây là package chứa các function handler
import (
	"myapp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// cấu trúc 1 handler luôn luôn nhận 1 e echo.Context và trả về 1 error
// viết hoa chữ cái đầu của function để export to anorther package. Ex: main
func Login(c echo.Context) error {
	return c.JSON(http.StatusOK, &models.LoginResponse{Token: "123456"})
}
