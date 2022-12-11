package handler

// mỗi một api sẽ gọi đến 1 function handler
// đây là package chứa các function handler
import (
	"log"
	"myapp/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// cấu trúc 1 handler luôn luôn nhận 1 e echo.Context và trả về 1 error
// viết hoa chữ cái đầu của function để export to anorther package. Ex: main
func Login(c echo.Context) error {
	// ta lấy username đã lưu ở middleware BasicAuth
	// tk c.Get này trả về interface, nên ta ép kiểu luôn thành string
	username := c.Get("username").(string)
	log.Println("Login username:", username)

	admin := c.Get("admin").(bool)
	log.Println("Login admin:", admin)

	// begin create jwt =============================================================
	// Set custom claims
	claims := &models.JwtCustomClaims{
		username,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		log.Println("Error create token")
		return err
	}
	// end create jwt =============================================================

	return c.JSON(http.StatusOK, &models.LoginResponse{Token: t})
}
