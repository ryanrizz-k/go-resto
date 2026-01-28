package user

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rynr00/go-resto/internal/model"
	"github.com/rynr00/go-resto/internal/tracing"
)

type Claims struct {
	jwt.StandardClaims
}

func (ur *userRepo) CreateUserSession(ctx context.Context, userID string) (model.UserSession, error) {
	ctx, span := tracing.CreateSpan(ctx, "CheckUserSession")
	defer span.End()

	accessToken, err := ur.generatedAccessToken(ctx, userID)
	if err != nil {
		return model.UserSession{}, err
	}

	return model.UserSession{
		JWTToken: accessToken,
	}, nil
}

func (ur *userRepo) generatedAccessToken(ctx context.Context, userID string) (string, error) {
	ctx, span := tracing.CreateSpan(ctx, "generatedAccessToken")
	defer span.End()

	accessTokenExp := time.Now().Add(ur.accessExp).Unix()
	accessClaims := Claims{
		jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: accessTokenExp,
		},
	}
	accessJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), accessClaims)

	return accessJwt.SignedString(ur.signKey)
}

func (ur *userRepo) CheckSession(ctx context.Context, data model.UserSession) (userID string, err error) {
	ctx, span := tracing.CreateSpan(ctx, "CheckSession")
	defer span.End()

	accessToken, err := jwt.ParseWithClaims(data.JWTToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return &ur.signKey.PublicKey, nil
	})
	if err != nil {
		return "", err
	}

	accessTokenClaims, ok := accessToken.Claims.(*Claims)
	if !ok {
		return "", errors.New("unauthorized")
	}

	if accessToken.Valid {
		return accessTokenClaims.Subject, nil
	}

	return "", errors.New("unauthorized")
}
