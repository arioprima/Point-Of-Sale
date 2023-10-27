package response

type LoginResponse struct {
	ID        string  `json:"user_id"`
	UserName  string  `json:"username"`
	UserRole  *string `json:"role"`
	TokenType string  `json:"token_type"`
	Token     string  `json:"token"`
}
