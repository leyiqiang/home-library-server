package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
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

	if len(statusCode) < 0 {
		WriteJSON(w, http.StatusBadRequest, theError, "error")
	}

	status := statusCode[0]
	WriteJSON(w, status, theError, "error")
}
