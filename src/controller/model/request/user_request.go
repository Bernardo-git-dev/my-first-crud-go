package request

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password" `
	Name     string `json:"name"`
	Age      int    `json:"age"` // Changed to string to match the response type
}
