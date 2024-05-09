package user

type UserRequest struct {
	Id       string `json:"id"`
	Name     string `validate:"required min=1,max=100" json:"name" form:"name"`
	Email    string `validate:"required min=1,max=150" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
