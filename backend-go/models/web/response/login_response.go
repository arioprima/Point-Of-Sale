package response

type LoginResponse struct {
	ID        string  `json:"id"`
	UserName  string  `json:"username"`
	Role      *string `json:"role"`
	TokenType string  `json:"token_type"`
	Token     string  `json:"token"`
}
