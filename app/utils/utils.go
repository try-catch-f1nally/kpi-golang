package utils

import (
	"encoding/json"
	"net/http"
	"os"
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
