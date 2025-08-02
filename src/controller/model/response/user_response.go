package response

type UserResponse struct {
	ID    string `json:"id",omitempty`
	Email string `json:"email" binding:"required,email"`
	Name  string `json:"name" binding:"required,min=3,max=50"`
	Age   int8   `json:"age" binding:"required,min=18,max=100"`
}
