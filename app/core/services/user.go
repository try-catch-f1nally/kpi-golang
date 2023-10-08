package services

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"kpi-golang/app/core"
	"kpi-golang/app/core/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (service *UserService) ChangeFirstName(userID uint, firstName string) error {
	return service.userRepository.UpdateFirstName(userID, firstName)
}

func (service *UserService) ChangeLastName(userID uint, lastName string) error {
	return service.userRepository.UpdateLastName(userID, lastName)
}

func (service *UserService) ChangeEmail(userID uint, email string) error {
	_, err := service.userRepository.GetByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err == nil {
		return &core.BadRequestError{Message: fmt.Sprintf("user with email %q already exists", email)}
	}
	return service.userRepository.UpdateEmail(userID, email)
}

func (service *UserService) ChangePassword(userID uint, oldPassword string, newPassword string) error {
	user, err := service.userRepository.Get(userID)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(oldPassword))
	if err != nil {
		return &core.BadRequestError{Message: "incorrect old password"}
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return service.userRepository.UpdatePassword(userID, passwordHash)
}
