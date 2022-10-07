package v1

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"testproject/controller/dto"
)

func RespondSuccess(w http.ResponseWriter, message string, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := dto.Response{
		Status:  true,
		Message: message,
		Payload: payload,
	}

	err := json.NewEncoder(w).Encode(resp)

	if err != nil {
		fmt.Println(err)
	}
}

func RespondError(w http.ResponseWriter, message string, payload interface{}, status int) dto.Response {
	w.Header().Add("Content-Type", "application/json")

	if status == 0 {
		status = http.StatusBadRequest
	}

	w.WriteHeader(status)

	resp := dto.Response{
		Message: message,
		Payload: payload,
	}

	err := json.NewEncoder(w).Encode(resp)

	if err != nil {
		fmt.Println(err)
	}
	return resp
}

func IsValidUUID(u string) bool {
	_, err := uuid.FromString(u)
	return err == nil
}
