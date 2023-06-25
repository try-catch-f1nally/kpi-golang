//go:build unit_test

package services

import (
	"errors"
	"kpi-golang/app/models"
	"testing"
)

func TestChangeFirstName(t *testing.T) {
	userGetError := errors.New("failed to get user")
	testCases := []struct {
		description     string
		userGetSuccess  *models.User
		userGetError    error
		expectedSuccess *models.User
		expectedError   error
	}{
		{
			description:   "get error from user repository",
			userGetError:  userGetError,
			expectedError: userGetError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			service := NewUserService(&UserRepositoryMock{GetError: tc.userGetError})
			err := service.ChangeFirstName(21, "Linus")
			if err != tc.expectedError {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedError, err)
			}
		})
	}
}

func TestChangeLastName(t *testing.T) {
	userGetError := errors.New("failed to get user")
	testCases := []struct {
		description     string
		userGetSuccess  *models.User
		userGetError    error
		expectedSuccess *models.User
		expectedError   error
	}{
		{
			description:   "get error from user repository",
			userGetError:  userGetError,
			expectedError: userGetError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			service := NewUserService(&UserRepositoryMock{GetError: tc.userGetError})
			err := service.ChangeLastName(21, "Torvald")
			if err != tc.expectedError {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedError, err)
			}
		})
	}
}

func TestChangeEmail(t *testing.T) {
	userGetError := errors.New("failed to get user")
	testCases := []struct {
		description     string
		userGetSuccess  *models.User
		userGetError    error
		expectedSuccess *models.User
		expectedError   error
	}{
		{
			description:   "get error from user repository",
			userGetError:  userGetError,
			expectedError: userGetError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			service := NewUserService(&UserRepositoryMock{GetError: tc.userGetError})
			err := service.ChangeEmail(21, "linus@example.com")
			if err != tc.expectedError {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedError, err)
			}
		})
	}
}

func TestChangePassword(t *testing.T) {
	userGetError := errors.New("failed to get user")
	testCases := []struct {
		description     string
		userGetSuccess  *models.User
		userGetError    error
		expectedSuccess *models.User
		expectedError   error
	}{
		{
			description:   "get error from user repository",
			userGetError:  userGetError,
			expectedError: userGetError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			service := NewUserService(&UserRepositoryMock{GetError: tc.userGetError})
			err := service.ChangePassword(21, "old password", "updated password")
			if err != tc.expectedError {
				t.Fatalf("%s. expected %v, got %v", tc.description, tc.expectedError, err)
			}
		})
	}
}
