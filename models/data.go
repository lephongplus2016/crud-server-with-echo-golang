// struct for response as result of post method
package models

type LoginResponse struct {
	Token string `json:"token"`
}

type X struct {
	Text string `json:"text"`
}
