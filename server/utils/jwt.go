package utils

import (
	"context"
	"fmt"
	"time"
	"userman/server/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

func CreateToken(ctx context.Context, username string, userUUID uuid.UUID, role string) (*string, error) {
	logger := logx.WithContext(ctx)
	logger.Infof("start create user %s token", username)
	claim := model.Claims{
		UUID:     userUUID,
		Username: username,
		Roles:    role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    model.Issuer,
			ExpiresAt: time.Now().Unix() + model.ExpiresTime,
		},
	}

	method := jwt.SigningMethodHS256
	token, err := jwt.NewWithClaims(method, claim).SignedString(model.Sign)
	if err != nil {
		logger.Errorf("create jwt token failed, err: %v", err)
		return nil, err
	}

	return &token, nil
}

func DecodeToken(ctx context.Context, token string) (*model.Claims, error) {
	logger := logx.WithContext(ctx)
	logger.Infof("start decode token %s", token)
	jwtInfo, err := jwt.ParseWithClaims(token, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return model.Sign, nil
	})

	if err != nil {
		logger.Errorf("parse token %s failed, err: %s", token, err)
		return nil, err
	}

	claim, ok := jwtInfo.Claims.(*model.Claims)
	if !ok {
		logger.Errorf("jwtinfo claims is vailed %v", jwtInfo.Claims)
		return nil, fmt.Errorf("invaild token")
	}

	return claim, nil
}
