package models

type User struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password *string `json:"password"`
	Roles    []struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"roles"`
}
