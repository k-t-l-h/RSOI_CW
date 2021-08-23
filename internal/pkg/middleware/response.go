package middleware

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"rsoi-kp-k-t-l-h/internal/models"
)

var DefaultError = models.Error{
	Message: "This is my error message",
}

func Response(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case models.StatusOkey:
		w.WriteHeader(http.StatusOK)
		if body != nil {
			jsn, _ := json.Marshal(body)
			_, _ = w.Write(jsn)
		}
	case models.StatusBadUUID:
		w.WriteHeader(http.StatusBadRequest)
		jsn, _ := json.Marshal(DefaultError)
		_, _ = w.Write(jsn)

	case models.StatusConflict:
		w.WriteHeader(http.StatusConflict)
		jsn, _ := json.Marshal(DefaultError)
		_, _ = w.Write(jsn)

	case models.StatusNotFound:
		w.WriteHeader(http.StatusNotFound)
		jsn, _ := json.Marshal(DefaultError)
		_, _ = w.Write(jsn)

	case models.StatusError:
		w.WriteHeader(http.StatusUnprocessableEntity)
		jsn, _ := json.Marshal(DefaultError)
		_, _ = w.Write(jsn)
	case models.StatusNoAuth:
		w.WriteHeader(http.StatusUnauthorized)
		jsn, _ := json.Marshal(DefaultError)
		_, _ = w.Write(jsn)
	default:
		w.WriteHeader(http.StatusTeapot)
	}

}

func UserUUID(r *http.Request) uuid.UUID {
	return uuid.UUID{}
}
