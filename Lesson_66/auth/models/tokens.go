package models

type RegisterRequest struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccessToken struct {
	Token string `json:"access_token"`
}

type RefreshToken struct {
	Token string `json:"refresh_token"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenDetails struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
	Expiry string `json:"expiry"`
}
