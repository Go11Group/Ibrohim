package models

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
