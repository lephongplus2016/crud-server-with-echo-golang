package handler

// mỗi một api sẽ gọi đến 1 function handler
// đây là package chứa các function handler
import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"myapp/models"
)

func User2(c echo.Context) error {

	return c.String(http.StatusOK, "user 2")
}

func User(c echo.Context) error {

	return c.String(http.StatusOK, "user 1")
}

var listUsers = []models.User{
	{
		Name: "Phong",
		Age:  22,
	},
	{
		Name: "Phú",
		Age:  22,
	},
	{
		Name: "Trung",
		Age:  22,
	},
	{
		Name: "Chính",
		Age:  22,
	},
	{
		Name: "Lộc",
		Age:  22,
	},
	{
		Name: "Nghĩa",
		Age:  22,
	},
	{
		Name: "Nhật",
		Age:  22,
	},
}

// biểu diễn streaming data
// mở bằng chrome dễ nhìn hơn postman
// 1 hàm trả về error chứ gì:
// thực ra nó trả về lỗi hoặc nil
func GetAllUsers(c echo.Context) error {
	// set header của response là "application-json"
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)

	// create an encoder
	// giống như tạo ra json để viết vào
	enc := json.NewEncoder(c.Response())

	// traverse throught the user list
	// _ là chỉ mục, ta dùng blank identifier để thông báo rằng biến này không bao giờ sử dụng.
	// -> tránh lỗi của go
	for _, user := range listUsers {
		// enc.Encode : lệnh để ghi vào encode
		if err := enc.Encode(user); err != nil {
			return err
		}

		// flush// đẩy, bắn  dữ liệu, theo từng item
		c.Response().Flush()
		// làm chậm lại để thấy data về  từ từ
		time.Sleep(1 * time.Second)
	}

	return nil
}
