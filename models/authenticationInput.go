package models

type AuthenticationInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
