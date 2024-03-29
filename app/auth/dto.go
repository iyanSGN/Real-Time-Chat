package auth

type UserResponseDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	MainPhoto string `json:"main_photo"`
}

type UserRequest struct {
	Name      string `form:"name"`
	Email     string `form:"email"`
	Phone     string `form:"phone"`
	Password  string `form:"password"`
	MainPhoto string `form:"main_photo"`
}