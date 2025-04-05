package dto

type User struct {
	Username string `json:"username" example:"username"`
	Email    string `json:"email" email:"username@gmail.com"`
}

type UserResponse struct {
	ID       uint   `json:"id" example:"1"`
	Username string `json:"username" example:"username"`
	Email    string `json:"email" example:"username@gmail.com"`
}
