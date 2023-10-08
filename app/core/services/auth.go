package services

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"kpi-golang/app/core"
	"kpi-golang/app/core/models"
	"kpi-golang/app/core/repositories"
)

type AuthService struct {
	userRepository repositories.UserRepository
	tokenService   *TokenService
}

func NewAuthService(userRepository repositories.UserRepository, tokenService *TokenService) *AuthService {
	return &AuthService{userRepository, tokenService}
}

type UserData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	UserId       uint   `json:"userId"`
}

type RegisterBody struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (service *AuthService) Register(registerBody *RegisterBody) (*UserData, error) {
	_, err := service.userRepository.GetByEmail(registerBody.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err == nil {
		return nil, &core.BadRequestError{
			Message: fmt.Sprintf("user with email %q already exists", registerBody.Email),
		}
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Email:     registerBody.Email,
		Password:  passwordHash,
		FirstName: registerBody.FirstName,
		LastName:  registerBody.LastName,
	}

	err = service.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	tokens, err := service.tokenService.GenerateTokens(user.ID)
	if err != nil {
		return nil, err
	}

	err = service.userRepository.UpdateToken(user.ID, tokens.refreshToken)
	if err != nil {
		return nil, err
	}

	return &UserData{
		AccessToken:  tokens.accessToken,
		RefreshToken: tokens.refreshToken,
		UserId:       user.ID,
	}, nil
}

func (service *AuthService) Login(loginBody *LoginBody) (*UserData, error) {
	user, err := service.userRepository.GetByEmail(loginBody.Email)
	wrongEmailOrPasswordError := &core.BadRequestError{Message: "wrong email or password"}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, wrongEmailOrPasswordError
	} else if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(loginBody.Password))
	if err != nil {
		return nil, wrongEmailOrPasswordError
	}

	tokens, err := service.tokenService.GenerateTokens(user.ID)
	if err != nil {
		return nil, err
	}

	err = service.userRepository.UpdateToken(user.ID, tokens.refreshToken)
	if err != nil {
		return nil, err
	}

	return &UserData{
		AccessToken:  tokens.accessToken,
		RefreshToken: tokens.refreshToken,
		UserId:       user.ID,
	}, nil
}

func (service *AuthService) Logout(token string) error {
	userID, err := service.tokenService.ValidateRefreshToken(token)
	if err != nil {
		return &core.BadRequestError{Message: "invalid refresh token provided"}
	}

	_, err = service.userRepository.Get(userID)
	if err != nil {
		return err
	}

	return service.userRepository.UpdateToken(userID, "")
}

func (service *AuthService) Refresh(token string) (*UserData, error) {
	userID, err := service.tokenService.ValidateRefreshToken(token)
	if err != nil {
		return nil, &core.BadRequestError{Message: "invalid refresh token provided"}
	}

	user, err := service.userRepository.Get(userID)
	if err != nil {
		return nil, err
	}

	currentToken := user.Token
	if currentToken != token {
		return nil, &core.BadRequestError{Message: "invalid refresh token provided"}
	}

	err = service.userRepository.UpdateToken(userID, "")
	if err != nil {
		return nil, err
	}

	tokens, err := service.tokenService.GenerateTokens(user.ID)
	if err != nil {
		return nil, err
	}

	err = service.userRepository.UpdateToken(userID, tokens.refreshToken)
	if err != nil {
		return nil, err
	}

	return &UserData{
		AccessToken:  tokens.accessToken,
		RefreshToken: tokens.refreshToken,
		UserId:       user.ID,
	}, nil
}
