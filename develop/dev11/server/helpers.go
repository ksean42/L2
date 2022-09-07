package server

import (
	"encoding/json"
	"net/http"
)

func (s *server) getError(text string) Error {
	return Error{
		Error: text,
	}
}

func (s *server) sendError(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(s.getError(err.Error()))
}
