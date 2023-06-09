package services

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"kpi-golang/app/models"
	"kpi-golang/app/utils"
)

type UserService struct {
	Db *gorm.DB
}

func (service *UserService) ChangeFirstName(userId uint, firstName string) error {
	return service.Db.Model(&models.User{}).Where("id = ?", userId).Update("first_name", firstName).Error
}

func (service *UserService) ChangeLastName(userId uint, lastName string) error {
	return service.Db.Model(&models.User{}).Where("id = ?", userId).Update("last_name", lastName).Error
}

func (service *UserService) ChangeEmail(userId uint, email string) error {
	var user models.User
	err := service.Db.Where("email = ?", email).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err == nil {
		return &utils.BadRequestError{Message: fmt.Sprintf("user with email %q already exists", email)}
	}
	return service.Db.First(&user, userId).Update("email", email).Error
}

func (service *UserService) ChangePassword(userId uint, oldPassword string, newPassword string) error {
	var user models.User
	err := service.Db.First(&user, userId).Error
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(oldPassword))
	if err != nil {
		return &utils.BadRequestError{Message: "incorrect old password"}
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	return service.Db.First(&user, userId).Update("password", passwordHash).Error
}
