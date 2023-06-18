package services

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"kpi-golang/app/models"
	"kpi-golang/app/utils"
)

type AuthService struct {
	userRepository UserRepository
	tokenService   *TokenService
}

func NewAuthService(userRepository UserRepository, tokenService *TokenService) *AuthService {
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
	_, err := service.userRepository.UserGetByEmail(registerBody.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err == nil {
		return nil, &utils.BadRequestError{
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

	err = service.userRepository.UserCreate(user)
	if err != nil {
		return nil, err
	}

	tokens, err := service.tokenService.GenerateTokens(user.ID)
	if err != nil {
		return nil, err
	}

	err = service.userRepository.UserUpdateToken(user.ID, tokens.refreshToken)
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
	user, err := service.userRepository.UserGetByEmail(loginBody.Email)
	wrongEmailOrPasswordError := &utils.BadRequestError{Message: "wrong email or password"}
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

	err = service.userRepository.UserUpdateToken(user.ID, tokens.refreshToken)
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
	userId, err := service.tokenService.ValidateRefreshToken(token)
	if err != nil {
		return &utils.BadRequestError{Message: "invalid refresh token provided"}
	}

	_, err = service.userRepository.UserGet(userId)
	if err != nil {
		return err
	}

	return service.userRepository.UserUpdateToken(userId, "")
}

func (service *AuthService) Refresh(token string) (*UserData, error) {
	userId, err := service.tokenService.ValidateRefreshToken(token)
	if err != nil {
		return nil, &utils.BadRequestError{Message: "invalid refresh token provided"}
	}

	user, err := service.userRepository.UserGet(userId)
	if err != nil {
		return nil, err
	}

	currentToken := user.Token
	if currentToken != token {
		return nil, &utils.BadRequestError{Message: "invalid refresh token provided"}
	}

	err = service.userRepository.UserUpdateToken(userId, "")
	if err != nil {
		return nil, err
	}

	tokens, err := service.tokenService.GenerateTokens(user.ID)
	if err != nil {
		return nil, err
	}

	err = service.userRepository.UserUpdateToken(userId, tokens.refreshToken)
	if err != nil {
		return nil, err
	}

	return &UserData{
		AccessToken:  tokens.accessToken,
		RefreshToken: tokens.refreshToken,
		UserId:       user.ID,
	}, nil
}
