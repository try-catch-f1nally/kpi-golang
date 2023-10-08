package repositories

import (
	"kpi-golang/app/core/models"
)

type UserRepository interface {
	Create(user *models.User) error
	Get(userID uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	UpdateFirstName(userID uint, firstName string) error
	UpdateLastName(userID uint, lastName string) error
	UpdateEmail(userID uint, email string) error
	UpdatePassword(userID uint, password []byte) error
	UpdateToken(userID uint, token string) error
}
