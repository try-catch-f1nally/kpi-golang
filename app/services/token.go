package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"kpi-golang/app/utils"
	"time"
)

type TokenService struct{}

type Tokens struct {
	accessToken  string
	refreshToken string
}

var accessTokenSecret = []byte(utils.GetEnvVar("ACCESS_TOKEN_SECRET", "access-token-secret"))
var refreshTokenSecret = []byte(utils.GetEnvVar("REFRESH_TOKEN_SECRET", "refresh-token-secret"))
var AccessTokenExp = time.Now().Add(time.Minute * 30).Unix()
var RefreshTokenExp = time.Now().Add(time.Hour * 24 * 30).Unix()

func (t *TokenService) GenerateTokens(userId uint) (*Tokens, error) {
	accessToken, err := generateToken(userId, accessTokenSecret, AccessTokenExp)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateToken(userId, refreshTokenSecret, RefreshTokenExp)
	if err != nil {
		return nil, err
	}

	return &Tokens{accessToken: accessToken, refreshToken: refreshToken}, nil
}

func generateToken(userId uint, secretKey []byte, exp int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"exp": exp, "sub": userId})
	return token.SignedString(secretKey)
}

func (t *TokenService) validateAccessToken(token string) (uint, error) {
	return validateToken(token, accessTokenSecret)
}

func (t *TokenService) validateRefreshToken(token string) (uint, error) {
	return validateToken(token, refreshTokenSecret)
}

func validateToken(token string, secretKey []byte) (uint, error) {
	invalidTokenError := errors.New("invalid token")
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidTokenError
		}
		return secretKey, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !parsedToken.Valid || !ok {
		return 0, invalidTokenError
	}

	userId, ok := claims["sub"].(float64)
	if !ok {
		return 0, invalidTokenError
	}
	return uint(userId), nil
}
