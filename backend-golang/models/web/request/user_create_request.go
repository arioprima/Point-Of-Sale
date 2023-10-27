package request

type UserCreateRequest struct {
	FirstName string  `json:"firstname" validate:"required,min=2,max=20"`
	LastName  *string `json:"lastname"`
	UserName  string  `json:"username" validate:"required,min=2,max=20"`
	Email     string  `json:"email" validate:"required,email"`
	Password  string  `json:"password" validate:"required,min=4,max=20"`
	UserRole  *string `json:"role"`
	UserImage *string `json:"image"`
}
