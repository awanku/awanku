package apihelper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func JSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Add("content-type", "application/json")

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		panic(fmt.Errorf("failed to marshal json: %v", err))
	}
	w.WriteHeader(status)
}

func BadRequestErrResp(w http.ResponseWriter, payload interface{}) {
	if _, ok := payload.(validation.Errors); ok {
		JSON(w, http.StatusBadRequest, map[string]interface{}{
			"errors": payload,
		})
		return
	}

	JSON(w, http.StatusBadRequest, payload)
}

func InternalServerErrResp(w http.ResponseWriter, err error) {
	log.Println("internal server err:", err)
	JSON(w, http.StatusInternalServerError, map[string]string{
		"error": "something's wrong on our side :(",
	})
}

func RedirectResp(w http.ResponseWriter, to string) {
	w.Header().Add("Location", to)
	w.WriteHeader(http.StatusFound)
}
