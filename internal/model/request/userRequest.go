package request

type UserRequest struct {
	Name     string `json:"name" binding:"required,min=4,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=50,contains=!@#$%&*"`
	Age      int8   `json:"age" binding:"required,min=2,max=127"`
}
