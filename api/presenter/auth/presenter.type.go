package presenter

type AuthResult struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	UserId       int    `json:"userId"`
	Username     string `json:"username"`
}

type AuthResp struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Results *AuthResult `json:"results"`
}
