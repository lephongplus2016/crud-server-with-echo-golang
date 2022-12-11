// struct for response as result of post method
package models

import "github.com/golang-jwt/jwt"

type LoginResponse struct {
	Token string `json:"token"`
}

type X struct {
	Text string `json:"text"`
}

type JwtCustomClaims struct {
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}
