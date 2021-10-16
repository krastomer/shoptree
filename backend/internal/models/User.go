package models

type User struct {
	ID       uint32 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"username"`
	Password string `json:"password"`
	Level    string `json:"level"`
}
