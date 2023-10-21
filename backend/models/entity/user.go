package entity

import "time"

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  *string   `json:"lastname"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      *string   `json:"role"`
	Image     *string   `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
