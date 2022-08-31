package handler

import (
	"net/http"

	"github.com/yt-lite/libs/errs"
)

func WriteResponse(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	b, ok := body.(*[]byte)

	if ok {
		w.Write(*b)
	} else {
		w.Write([]byte(body.(string)))
	}
}

func WriteError(w http.ResponseWriter, err *errs.AppError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	w.Write([]byte(err.Error()))
}

