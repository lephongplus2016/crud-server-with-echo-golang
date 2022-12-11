package mdw

import "github.com/labstack/echo/v4"

// signature method of middleware
func BasicAuth(username string, password string, c echo.Context) (bool, error) {
	if username == "admin" || password == "123" {
		// cách tạo ra 1 global variable, sử dụng được ở mọi nơi
		// thực ra lấy đc ở mọi nơi có c echo.Context
		c.Set("username", username)
		c.Set("admin", true)
		return true, nil
	}

	if username == "phong" || password == "123" {
		// cách tạo ra 1 global variable, sử dụng được ở mọi nơi
		// thực ra lấy đc ở mọi nơi có c echo.Context
		c.Set("username", username)
		c.Set("admin", false)
		return true, nil
	}
	return false, nil
}
