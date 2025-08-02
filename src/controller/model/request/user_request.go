package request

type UserRequest struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Age      string `json:"name" binding:"required,min=3,max=50"`
}
