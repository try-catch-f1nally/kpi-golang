package services

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"kpi-golang/app/models"
	"kpi-golang/app/utils"
)

type UserRepository interface {
	UserCreate(user *models.User) error
	UserGet(userId uint) (*models.User, error)
	UserGetByEmail(email string) (*models.User, error)
	UserUpdateFirstName(userId uint, firstName string) error
	UserUpdateLastName(userId uint, lastName string) error
	UserUpdateEmail(userId uint, email string) error
	UserUpdatePassword(userId uint, password []byte) error
	UserUpdateToken(userId uint, token string) error
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{userRepository}
}

func (service *UserService) ChangeFirstName(userId uint, firstName string) error {
	return service.userRepository.UserUpdateFirstName(userId, firstName)
}

func (service *UserService) ChangeLastName(userId uint, lastName string) error {
	return service.userRepository.UserUpdateLastName(userId, lastName)
}

func (service *UserService) ChangeEmail(userId uint, email string) error {
	_, err := service.userRepository.UserGetByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err == nil {
		return &utils.BadRequestError{Message: fmt.Sprintf("user with email %q already exists", email)}
	}
	return service.userRepository.UserUpdateEmail(userId, email)
}

func (service *UserService) ChangePassword(userId uint, oldPassword string, newPassword string) error {
	user, err := service.userRepository.UserGet(userId)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(oldPassword))
	if err != nil {
		return &utils.BadRequestError{Message: "incorrect old password"}
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return service.userRepository.UserUpdatePassword(userId, passwordHash)
}
