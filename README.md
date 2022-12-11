# crud-server-with-echo-golang

## Document

-   Echo https://echo.labstack.com/

-   JWT https://echo.labstack.com/cookbook/jwt/

## Note

-   Tk golang nó tạo struct để dùng thay cho json object.

-   Cả body của request và body của response đều dùng struct cả.

-   Không được đặt folder, package trùng tên key word của go như: `middleware`, `server`, `echo`

### If in golang

-   https://www.callicoder.com/golang-control-flow/#if-with-a-short-statement
-   If with a short statement:

-   An `if` statement in Golang can also contain a short declaration statement preceding the conditional expression

-   ```Go
    if n := 10; n%2 == 0 {
    fmt.Printf("%d is even\n", n)
    }

    ```

-   The variable declared in the short statement is only available inside the if block and it’s else or else-if branches

-   ```Go
    if n := 15; n%2 == 0 {
        fmt.Printf("%d is even\n", n)
    } else {
        fmt.Printf("%d is odd\n", n)
    }

    ```

### UnMarshaling or consuming Json Data

-   The `Unmarshal` function provided by Go’s JSON standard library lets us parse `raw JSON data` in the form of `[]byte` variables.

### Marshaling or creating Json Data

-   In above code, `Json.Marshal()` return , finalJson as the Json string, represented as bytes, and second parameter here is the error Which should be ideally handle.

### Encoding JSON : writing | creating data

### Decoding JSON | Reading data

## Flow chart

1. User gọi api "/login" để đăng nhập, nó sẽ trả về cho user 1 tk json web token.

2. User sử dụng jwt đó để goi api "/", nếu jwt hợp lệ thì mới cho đi qua

![Flow chart](https://github.com/lephongplus2016/crud-server-with-echo-golang/blob/master/img/jwt_diagram.png?raw=true)

