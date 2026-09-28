package requests

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Age      int8   `json:"age"`
}
