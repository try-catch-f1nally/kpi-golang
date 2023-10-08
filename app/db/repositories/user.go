package repositories

import (
	"gorm.io/gorm"
	"kpi-golang/app/core/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) Create(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *UserRepository) Get(userID uint) (*models.User, error) {
	var user models.User
	err := repo.db.First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) UpdateFirstName(userID uint, firstName string) error {
	return repo.db.Model(&models.User{}).Where("id = ?", userID).Update("first_name", firstName).Error
}

func (repo *UserRepository) UpdateLastName(userID uint, lastName string) error {
	return repo.db.Model(&models.User{}).Where("id = ?", userID).Update("last_name", lastName).Error
}

func (repo *UserRepository) UpdateEmail(userID uint, email string) error {
	return repo.db.Model(&models.User{}).Where("id = ?", userID).Update("email", email).Error
}

func (repo *UserRepository) UpdatePassword(userID uint, password []byte) error {
	return repo.db.Model(&models.User{}).Where("id = ?", userID).Update("password", password).Error
}

func (repo *UserRepository) UpdateToken(userID uint, token string) error {
	return repo.db.Model(&models.User{}).Where("id = ?", userID).Update("token", token).Error
}
