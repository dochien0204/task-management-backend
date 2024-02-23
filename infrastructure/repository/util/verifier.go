package util

import (
	"context"
	"encoding/json"
	"fmt"
	"source-base-go/entity"
	"time"

	"github.com/go-redis/redis/v8"
)

type UserData struct {
	Status   string
	ListRole []string
}

type Verifier interface {
	VerifyToken(token string) (bool, *UserData, error)
	InvalidateToken(token string) error
	CacheUserData(user *entity.User, listRole []string, expiredAt int) error
	GetUserData(userId int) (*UserData, error)
}

type JWTVerifier struct {
	redis *redis.Client
}

func NewJWTVerifier(redis *redis.Client) *JWTVerifier {
	return &JWTVerifier{
		redis: redis,
	}
}

func (v *JWTVerifier) VerifyToken(token string) (bool, *UserData, error) {
	claims, err := ParseAccessToken(token)
	if err != nil {
		return false, nil, entity.ErrUnauthorized
	}

	value, err := v.redis.Get(context.TODO(), claims.Jti).Result()
	if err != redis.Nil || value == token {
		return false, nil, entity.ErrUnauthorized
	}

	userData, err := v.GetUserData(claims.UserId)
	if err != nil {
		return false, nil, entity.ErrUnauthorized
	}

	return true, userData, nil

}

func (v *JWTVerifier) InvalidateToken(token string) error {
	claims, err := ParseAccessToken(token)
	if err != nil {
		return err
	}

	err = v.redis.Set(context.TODO(), claims.Jti, token, time.Duration(time.Unix(claims.ExpiresAt, 0).Sub(time.Now()))).Err()
	if err != nil {
		return err
	}

	return nil
}

func (v *JWTVerifier) CacheUserData(user *entity.User, listRole []string, expiredAt int) error {
	if user.Status == nil {
		return entity.ErrInternalServerError
	}

	userData := UserData{
		Status:   user.Status.Code,
		ListRole: listRole,
	}

	userDataJSON, err := json.Marshal(&userData)
	if err != nil {
		return err
	}

	err = v.redis.Set(context.TODO(), fmt.Sprint(user.Id), userDataJSON, time.Duration(expiredAt)).Err()
	if err != nil {
		return err
	}

	return nil
}

func (v *JWTVerifier) GetUserData(userId int) (*UserData, error) {
	result, err := v.redis.Get(context.TODO(), fmt.Sprint(userId)).Result()
	if err != nil {
		return nil, entity.ErrUnauthorized
	}

	userData := UserData{}
	err = json.Unmarshal([]byte(result), &userData)
	if err != nil {
		return nil, entity.ErrLock
	}

	return &userData, nil
}
