package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yt-lite/libs/errs"
)

func WriteResponse(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := map[string]interface{}{
		"status": status,
		"data":   body,
	}

	respByte, _ := json.Marshal(resp)

	w.Write(respByte)

}

func WriteError(w http.ResponseWriter, err *errs.AppError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)

	resp := map[string]interface{}{
		"status": err.Code,
		"data": map[string]interface{}{
			"message": err.Message,
		},
	}

	respByte, _ := json.Marshal(resp)

	w.Write(respByte)

}
