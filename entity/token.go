package entity

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	jwt.StandardClaims
	UserId int    `json:"userId"`
	Jti    string `json:"jti"`
	Role string `json:"role"`
}

type RefreshToken struct {
	jwt.StandardClaims
	UserId int    `json:"userId"`
	Jti    string `json:"jti"`
}

type TokenPair struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
