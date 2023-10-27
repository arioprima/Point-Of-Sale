package response

import "time"

type UserResponse struct {
	ID        string    `json:"user_id"`
	FirstName string    `json:"firstname"`
	LastName  *string   `json:"lastname"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	UserRole  *string   `json:"role"`
	UserImage *string   `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
