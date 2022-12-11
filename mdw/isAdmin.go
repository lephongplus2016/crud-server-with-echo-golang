package mdw

import (
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// custome middleware là 1 handler function
// content: if admin in jwt === true, we handle continue
// else, we return error
func IsAdminMdw(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)

		claims := user.Claims.(jwt.MapClaims)

		admin := claims["admin"].(bool)
		log.Printf("isAdminMdw: %v\n", admin)

		if admin {
			next(c)
		}

		return echo.ErrUnauthorized
	}
}
