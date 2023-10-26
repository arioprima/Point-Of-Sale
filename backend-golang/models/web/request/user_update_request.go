package request

type UserUpdateRequest struct {
	ID        string  `json:"user_id" validate:"required"`
	FirstName string  `json:"firstname" validate:"required,min=2,max=20"`
	LastName  *string `json:"lastname"`
	Image     *string `json:"image"`
}
