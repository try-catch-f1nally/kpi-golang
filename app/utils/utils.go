package utils

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type UnauthorizedError struct {
	message string
}

type BadRequestError struct {
	Message string
}

func (e UnauthorizedError) Error() string {
	return e.message
}

func (e BadRequestError) Error() string {
	return e.Message
}

func HandleError(w http.ResponseWriter, err error) {
	var statusCode int
	switch err.(type) {
	case UnauthorizedError:
		statusCode = http.StatusUnauthorized
	case BadRequestError:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}
	http.Error(w, err.Error(), statusCode)
}

func GetEnvVar(name string, defaultValue string) string {
	env, envExists := os.LookupEnv(name)
	if !envExists {
		env = defaultValue
	}
	return env
}

func RespondWithJson(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		HandleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		HandleError(w, err)
	}
}

func AuthenticateRequest(r *http.Request, validateToken func(string) (uint, error)) (uint, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, UnauthorizedError{message: "authentication failed due to missing access token"}
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return 0, UnauthorizedError{message: "authentication failed due to missing access token"}
	}

	userId, err := validateToken(tokenParts[1])
	if err != nil {
		return 0, UnauthorizedError{message: "authentication failed due to invalid access token"}
	}

	return userId, nil
}
