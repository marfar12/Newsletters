package util

import (
	"encoding/json"
	"net/http"
)

type ErrorStruct struct {
	Error string
}

func WriteErrResponse(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	if err != nil {
		b, _ := json.Marshal(ErrorStruct{Error: err.Error()})
		w.Write(b)
	}
}

func WriteResponse(w http.ResponseWriter, statusCode int, body any) {
	w.WriteHeader(statusCode)
	b, _ := json.Marshal(body)
	w.Write(b)
}
