package utils

import (
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

//var ErrNotFound = errors.New("not found")

func GetHash(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", errors.New("error Encrypting string")
	}
	return string(hash), nil
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}, wrap ...string) error {
	wrapper := make(map[string]interface{}, 0)
	var js []byte
	var err error
	if len(wrap) > 0 {
		wrapper[wrap[0]] = data
		js, err = json.Marshal(wrapper)
	} else {
		js, err = json.Marshal(data)
	}

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)

	return err
}

func ErrorJSON(w http.ResponseWriter, err error, statusCode ...int) {
	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	if len(statusCode) <= 0 {
		WriteJSON(w, http.StatusBadRequest, theError)
		return
	}

	status := statusCode[0]
	WriteJSON(w, status, theError)
}
