//go:build integration_test

package repositories

import (
	"bytes"
	"gorm.io/gorm"
	"kpi-golang/app/db"
	"kpi-golang/app/models"
	"log"
	"testing"
)

func cleanUserTable(db *gorm.DB) {
	err := db.Exec("TRUNCATE TABLE users").Error
	if err != nil {
		log.Fatal(err)
	}
}

func TestUserCreate(t *testing.T) {
	database := db.Init()
	cleanUserTable(database)
	repo := NewUserRepository(database)

	newUser := &models.User{
		Email:     "email@example.com",
		Password:  []byte("password"),
		FirstName: "Linus",
		LastName:  "Torvalds",
	}

	t.Run("create and get new user", func(t *testing.T) {
		err := repo.UserCreate(newUser)
		if err != nil {
			t.Errorf("failed to create user with error: %v", err)
			return
		}

		user, err := repo.UserGet(newUser.ID)
		if err != nil {
			t.Errorf("failed to get user with error: %v", err)
			return
		}

		if newUser.ID != user.ID || newUser.Email != user.Email || !bytes.Equal(newUser.Password, user.Password) ||
			newUser.FirstName != user.FirstName || newUser.LastName != user.LastName {
			t.Errorf("user data is corrupted; actual: %v, expected: %v", user, newUser)
		}
	})
}

func TestUserGetByEmail(t *testing.T) {
	database := db.Init()
	cleanUserTable(database)
	repo := NewUserRepository(database)

	newUser := &models.User{
		Email:     "email@example.com",
		Password:  []byte("password"),
		FirstName: "Linus",
		LastName:  "Torvalds",
	}

	t.Run("create and get by email new user", func(t *testing.T) {
		err := repo.UserCreate(newUser)
		if err != nil {
			t.Errorf("failed to create user with error: %v", err)
			return
		}

		user, err := repo.UserGetByEmail(newUser.Email)
		if err != nil {
			t.Errorf("failed to get user with error: %v", err)
			return
		}

		if newUser.ID != user.ID || !bytes.Equal(newUser.Password, user.Password) ||
			newUser.FirstName != user.FirstName || newUser.LastName != user.LastName {
			t.Errorf("user data is corrupted; actual: %v, expected: %v", user, newUser)
		}
	})
}

func TestUserUpdateEmail(t *testing.T) {
	database := db.Init()
	cleanUserTable(database)
	repo := NewUserRepository(database)

	newUser := &models.User{
		Email:     "email@example.com",
		Password:  []byte("password"),
		FirstName: "Linus",
		LastName:  "Torvalds",
	}

	t.Run("create, update email and get user", func(t *testing.T) {
		err := repo.UserCreate(newUser)
		if err != nil {
			t.Errorf("failed to create user with error: %v", err)
			return
		}

		newEmail := "linus@example.com"
		err = repo.UserUpdateEmail(newUser.ID, newEmail)
		if err != nil {
			t.Errorf("failed to update user email with error: %v", err)
			return
		}

		user, err := repo.UserGet(newUser.ID)
		if err != nil {
			t.Errorf("failed to get user with error: %v", err)
			return
		}

		if user.Email != newEmail {
			t.Errorf("user email does not match updated one; actual: %v, expected: %v", user.Email, newEmail)
		}
	})
}

func TestUserUpdatePassword(t *testing.T) {
	database := db.Init()
	cleanUserTable(database)
	repo := NewUserRepository(database)

	newUser := &models.User{
		Email:     "email@example.com",
		Password:  []byte("password"),
		FirstName: "Linus",
		LastName:  "Torvalds",
	}

	t.Run("create, update password and get user", func(t *testing.T) {
		err := repo.UserCreate(newUser)
		if err != nil {
			t.Errorf("failed to create user with error: %v", err)
			return
		}

		newPassword := []byte("updated password")
		err = repo.UserUpdatePassword(newUser.ID, newPassword)
		if err != nil {
			t.Errorf("failed to update user password with error: %v", err)
			return
		}

		user, err := repo.UserGet(newUser.ID)
		if err != nil {
			t.Errorf("failed to get user with error: %v", err)
			return
		}

		if !bytes.Equal(user.Password, newPassword) {
			t.Errorf("user password does not match updated one; actual: %v, expected: %v", user.Password, newPassword)
		}
	})
}

func TestUserUpdateFirstName(t *testing.T) {
	database := db.Init()
	cleanUserTable(database)
	repo := NewUserRepository(database)

	newUser := &models.User{
		Email:     "email@example.com",
		Password:  []byte("password"),
		FirstName: "Linus",
		LastName:  "Torvalds",
	}

	t.Run("create, update first name and get user", func(t *testing.T) {
		err := repo.UserCreate(newUser)
		if err != nil {
			t.Errorf("failed to create user with error: %v", err)
			return
		}

		newFirstName := "Bill"
		err = repo.UserUpdateFirstName(newUser.ID, newFirstName)
		if err != nil {
			t.Errorf("failed to update user first name with error: %v", err)
			return
		}

		user, err := repo.UserGet(newUser.ID)
		if err != nil {
			t.Errorf("failed to get user with error: %v", err)
			return
		}

		if user.FirstName != newFirstName {
			t.Errorf("user first name does not match updated one; actual: %v, expected: %v", user.FirstName, newFirstName)
		}
	})
}

func TestUserUpdateLastName(t *testing.T) {
	database := db.Init()
	cleanUserTable(database)
	repo := NewUserRepository(database)

	newUser := &models.User{
		Email:     "email@example.com",
		Password:  []byte("password"),
		FirstName: "Linus",
		LastName:  "Torvalds",
	}

	t.Run("create, update last name and get user", func(t *testing.T) {
		err := repo.UserCreate(newUser)
		if err != nil {
			t.Errorf("failed to create user with error: %v", err)
			return
		}

		newLastName := "Gates"
		err = repo.UserUpdateFirstName(newUser.ID, newLastName)
		if err != nil {
			t.Errorf("failed to update user last name with error: %v", err)
			return
		}

		user, err := repo.UserGet(newUser.ID)
		if err != nil {
			t.Errorf("failed to get user with error: %v", err)
			return
		}

		if user.FirstName != newLastName {
			t.Errorf("user last name does not match updated one; actual: %v, expected: %v", user.LastName, newLastName)
		}
	})
}

func TestUserUpdateToken(t *testing.T) {
	database := db.Init()
	cleanUserTable(database)
	repo := NewUserRepository(database)

	newUser := &models.User{
		Email:     "email@example.com",
		Password:  []byte("password"),
		FirstName: "Linus",
		LastName:  "Torvalds",
		Token:     "token",
	}

	t.Run("create, update token and get user", func(t *testing.T) {
		err := repo.UserCreate(newUser)
		if err != nil {
			t.Errorf("failed to create user with error: %v", err)
			return
		}

		newToken := "updated token"
		err = repo.UserUpdateFirstName(newUser.ID, newToken)
		if err != nil {
			t.Errorf("failed to update user token with error: %v", err)
			return
		}

		user, err := repo.UserGet(newUser.ID)
		if err != nil {
			t.Errorf("failed to get user with error: %v", err)
			return
		}

		if user.FirstName != newToken {
			t.Errorf("user token does not match updated one; actual: %v, expected: %v", user.Token, newToken)
		}
	})
}
