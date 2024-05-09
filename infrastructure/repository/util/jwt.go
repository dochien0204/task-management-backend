package util

import (
	"errors"
	"source-base-go/config"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/define"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func keyFunc(key string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New(config.UNAUTHORIZED)
		}

		return []byte(key), nil
	}
}

func GetToken(ctx *gin.Context) (string, error) {
	//Get Authorization Header
	authHeader := ctx.GetHeader(config.GetString("jwt.header"))
	if len(authHeader) <= len(config.GetString("jwt.schema")) {
		return "", errors.New(config.UNAUTHORIZED)
	}

	//Get Token from Authorization Header ("jwt.schema" = Bearer)
	tokenString := authHeader[len(config.GetString("jwt.schema"))+1:]

	return tokenString, nil
}

// Validate Token
func ValidateAccessToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, keyFunc(config.GetString("jwt.secretKey")))
}

func ParseAccessToken(token string) (*entity.TokenClaims, error) {
	claims := entity.TokenClaims{}

	_, err := jwt.ParseWithClaims(token, &claims, keyFunc(config.GetString("jwt.secretKey")))
	if err != nil {
		return nil, err
	}

	return &claims, nil
}

func GenerateAccessToken(user *entity.User) (string, error) {
	//Get random
	random, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	var roleUser string
	for _, role := range user.Role {
		if role.Type == string(define.SYSTEM) {
			roleUser = role.Code
			break
		} 
	}

	//Create claims
	claims := entity.TokenClaims{
		UserId: user.Id,
		Jti:    random.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.GetInt("jwt.accessMaxAge"))).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Role: roleUser,
	}

	//create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//Sign with secret key
	signedToken, err := token.SignedString([]byte(config.GetString("jwt.secretKey")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Generate Refresh Token
func GenerateRefreshToken(user *entity.User) (string, error) {
	//Create claims
	claims := entity.RefreshToken{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.GetInt("jwt.refreshMaxAge"))).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	//Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//Sign with secret key
	signedToken, err := token.SignedString([]byte(config.GetString("jwt.secretKey")))
	if err != nil {
		return "", nil
	}

	return signedToken, nil
}
