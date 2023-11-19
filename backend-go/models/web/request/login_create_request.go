package request

type UserLoginRequest struct {
	UserName string `json:"username" validate:"required,min=2,max=20"`
	Password string `json:"password" validate:"required,min=4,max=20"`
}
