# crud-server-with-echo-golang

## Document

-   Echo https://echo.labstack.com/

-   JWT https://echo.labstack.com/cookbook/jwt/

## Note

-   Tk golang nó tạo struct để dùng thay cho json object.

-   Cả body của request và body của response đều dùng struct cả.

-   Không được đặt folder, package trùng tên key word của go như: `middleware`, `server`, `echo`

## Flow chart

1. User gọi api "/login" để đăng nhập, nó sẽ trả về cho user 1 tk json web token.

2. User sử dụng jwt đó để goi api "/", nếu jwt hợp lệ thì mới cho đi qua
