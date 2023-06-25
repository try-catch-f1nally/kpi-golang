package services

import "kpi-golang/app/models"

type UserRepositoryMock struct {
	GetSuccess  *models.User
	GetError    error
	CreateError error
}

func (repo *UserRepositoryMock) UserCreate(user *models.User) error {
	return repo.CreateError
}

func (repo *UserRepositoryMock) UserGet(userId uint) (*models.User, error) {
	return repo.GetSuccess, repo.GetError
}

func (repo *UserRepositoryMock) UserGetByEmail(email string) (*models.User, error) {
	return repo.GetSuccess, repo.GetError
}

func (repo *UserRepositoryMock) UserUpdateFirstName(userId uint, firstName string) error {
	return repo.GetError
}

func (repo *UserRepositoryMock) UserUpdateLastName(userId uint, lastName string) error {
	return repo.GetError
}

func (repo *UserRepositoryMock) UserUpdateEmail(userId uint, email string) error {
	return repo.GetError
}

func (repo *UserRepositoryMock) UserUpdatePassword(userId uint, password []byte) error {
	return repo.GetError
}

func (repo *UserRepositoryMock) UserUpdateToken(userId uint, token string) error {
	return repo.GetError
}
