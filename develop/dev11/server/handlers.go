package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"server/model"
	"strconv"
)

// Error структура ошибки
type Error struct {
	Error string `json:"error"`
}

// Result структура успешного ответа
type Result struct {
	Result interface{} `json:"result"`
}

func (s *server) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		s.sendError(w, errors.New("invalid request"), 503)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	decoded := model.Request{}
	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		s.sendError(w, err, http.StatusBadRequest)
		return
	}
	errCreate := s.storage.Create(decoded)
	if errCreate != nil {
		s.sendError(w, errCreate, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(Result{"OK"})
}

func (s *server) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	decoded := model.Request{}
	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		s.sendError(w, err, http.StatusBadRequest)
		return
	}
	err = s.storage.Update(decoded)
	if err != nil {
		s.sendError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(Result{"OK"})
}

func (s *server) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	decoded := model.Request{}
	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		s.sendError(w, err, http.StatusBadRequest)
		return
	}
	err = s.storage.Delete(decoded)
	if err != nil {
		s.sendError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(Result{"OK"})
}

func (s *server) eventsDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	args := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")
	userID, parseEr := strconv.Atoi(args.Get("user_id"))
	day := args.Get("day")
	ev, err := s.storage.GetByDay(userID, day)
	if parseEr != nil || day == "" {
		s.sendError(w, errors.New("bad input"), http.StatusBadRequest)
	} else if err != nil {
		s.sendError(w, err, http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(Result{ev})
	}
}

func (s *server) eventsWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	args := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")
	userID, parseEr := strconv.Atoi(args.Get("user_id"))
	day := args.Get("day")
	ev, err := s.storage.GetByWeek(userID, day)
	if parseEr != nil || day == "" {
		s.sendError(w, errors.New("bad input"), http.StatusBadRequest)
	} else if err != nil {
		s.sendError(w, err, http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(Result{ev})
	}
}

func (s *server) eventsMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	args := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")
	userID, parseEr := strconv.Atoi(args.Get("user_id"))
	day := args.Get("day")
	ev, err := s.storage.GetByMonth(userID, day)
	if parseEr != nil || day == "" {
		s.sendError(w, errors.New("bad input"), http.StatusBadRequest)
	} else if err != nil {
		s.sendError(w, err, http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(Result{ev})
	}
}
