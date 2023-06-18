package repositories

import (
	"gorm.io/gorm"
	"kpi-golang/app/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) UserCreate(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *UserRepository) UserGet(userId uint) (*models.User, error) {
	var user models.User
	err := repo.db.First(&user, userId).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) UserGetByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) UserUpdateFirstName(userId uint, firstName string) error {
	return repo.db.Model(&models.User{}).Where("id = ?", userId).Update("first_name", firstName).Error
}

func (repo *UserRepository) UserUpdateLastName(userId uint, lastName string) error {
	return repo.db.Model(&models.User{}).Where("id = ?", userId).Update("last_name", lastName).Error
}

func (repo *UserRepository) UserUpdateEmail(userId uint, email string) error {
	return repo.db.Model(&models.User{}).Where("id = ?", userId).Update("email", email).Error
}

func (repo *UserRepository) UserUpdatePassword(userId uint, password []byte) error {
	return repo.db.Model(&models.User{}).Where("id = ?", userId).Update("password", password).Error
}

func (repo *UserRepository) UserUpdateToken(userId uint, token string) error {
	return repo.db.Model(&models.User{}).Where("id = ?", userId).Update("token", token).Error
}
