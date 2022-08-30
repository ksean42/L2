package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"server/model"
	"strconv"
	"strings"
	"sync"
	"time"
)

type storage struct {
	sync.RWMutex
	events map[int][]model.Event
}

type Log struct {
	w io.Writer
}

func (l *Log) log(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer handler(w, r)
		body, err := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		data := strings.Builder{}
		data.WriteString(fmt.Sprintf("Request time: %s\nMethod: %s\n", time.Now(), "GET"))
		data.WriteString("\nHeader: \n")
		for k, v := range r.Header {
			data.WriteString(fmt.Sprintf("%s : ", k))
			for _, s := range v {
				data.WriteString(fmt.Sprintf("%s", s))
			}
			data.WriteString("\n")
		}
		data.WriteString("\nQuery params : \n")
		err = r.ParseForm()
		if err == nil {
			for k, v := range r.Form {
				data.WriteString(fmt.Sprintf("%s : ", k))
				for _, s := range v {
					data.WriteString(fmt.Sprintf("%s", s))
				}
				data.WriteString("\n")
			}
		} else {
			data.WriteString(err.Error() + "\n")
		}
		data.WriteString(fmt.Sprintf("Body:\n%s\n", string(body)))
		l.w.Write([]byte(data.String()))
	}

}

type server struct {
	storage storage
}
type Error struct {
	Error string `json:"error"`
}

type Result struct {
	Result interface{} `json:"result"`
}

func (s *server) getError(text string) Error {
	return Error{
		Error: text,
	}
}

func parseConfig() string {
	file, err := os.Open("config")
	if err != nil {
		log.Fatal(err)
	}
	read := bufio.NewScanner(file)
	for read.Scan() {
		str := strings.Split(read.Text(), "=")
		if str[0] == "port" {
			_, err := strconv.Atoi(str[1])
			if err != nil {
				log.Fatal("Incorrect config")
			}
			return str[1]
		}

	}
	return ""
}

func main() {
	port := parseConfig()
	logger := &Log{os.Stdout}
	storage := storage{events: make(map[int][]model.Event, 10)}
	s := &server{storage: storage}
	http.HandleFunc("/create_event", logger.log(s.createEvent))
	http.HandleFunc("/update_event", logger.log(s.updateEvent))
	http.HandleFunc("/delete_event", logger.log(s.deleteEvent))
	http.HandleFunc("/events_for_day", logger.log(s.eventsDay))
	http.HandleFunc("/events_for_week", logger.log(s.eventsWeek))
	http.HandleFunc("/events_for_month", logger.log(s.eventsMonth))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *server) parseTime(date string) (time.Time, error) {

	res, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, err
	}
	return res, nil
}
func (s *server) sendError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(s.getError(err.Error()))
}

func (s *server) addEvent(event model.Event) {
	s.storage.RWMutex.Lock()
	s.storage.events[event.UserId] = append(s.storage.events[event.UserId], event)

	s.storage.RWMutex.Unlock()
}

func (s *server) deleteEventStorage(event model.Event) error {
	s.storage.RWMutex.Lock()
	defer s.storage.RWMutex.Unlock()
	if _, ok := s.storage.events[event.UserId]; !ok {
		return errors.New("user doesnt exist")
	} else {
		e := s.storage.events[event.UserId]
		for i := 0; i < len(e); i++ {
			if e[i].Date == event.Date && e[i].Title == event.Title {
				e = append(e[:i], e[i+1:]...)
				s.storage.events[event.UserId] = e
				return nil
			}
		}
	}
	return errors.New("event doesnt exist")
}

func (s *server) updateEventStorage(event model.Event) error {
	s.storage.RWMutex.Lock()
	defer s.storage.RWMutex.Unlock()
	if _, ok := s.storage.events[event.UserId]; !ok {
		return errors.New("user doesnt exist")
	} else {
		e := s.storage.events[event.UserId]
		for i := 0; i < len(e); i++ {
			if e[i].Date == event.Date && e[i].Title == event.Title {
				e[i] = event
				return nil
			}
		}
	}
	return errors.New("event doesnt exist")
}

func (s *server) validateReq(r *http.Request) (model.Event, error) {
	decoded := model.Request{}
	err := json.NewDecoder(r.Body).Decode(&decoded)
	t, errTime := s.parseTime(decoded.Date)
	if err != nil {
		return model.Event{}, err
	}
	if errTime != nil {
		return model.Event{}, errTime
	}
	if decoded.Title == "" || decoded.UserId == 0 || decoded.Description == "" {
		return model.Event{}, errors.New("bad request")
	}
	newEvent := model.Event{
		UserId:      decoded.UserId,
		Date:        t,
		Title:       decoded.Title,
		Description: decoded.Description}
	return newEvent, nil

}
func (s *server) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	newEvent, err := s.validateReq(r)
	if err != nil {
		s.sendError(w, err)
		return
	}
	s.addEvent(newEvent)
	json.NewEncoder(w).Encode(Result{newEvent})
}

func (s *server) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	newEvent, err := s.validateReq(r)
	if err != nil {
		s.sendError(w, err)
		return
	}
	err = s.updateEventStorage(newEvent)
	if err != nil {
		s.sendError(w, err)
		return
	}
	json.NewEncoder(w).Encode(Result{newEvent})
}

func (s *server) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	newEvent, err := s.validateReq(r)
	if err != nil {
		s.sendError(w, err)
		return
	}
	err = s.deleteEventStorage(newEvent)
	if err != nil {
		s.sendError(w, err)
		return
	}
	json.NewEncoder(w).Encode(Result{"event deleted"})
}

func (s *server) getByDay(id int, day time.Time) ([]model.Event, error) {
	events, ok := s.storage.events[id]
	if !ok {
		return nil, errors.New("user doesnt exist")
	}
	res := make([]model.Event, 0, 10)
	for _, v := range events {
		if v.Date == day {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		return nil, errors.New("there are no events for this day")
	} else {
		return res, nil
	}
}

func (s *server) eventsDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	args := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")
	userId, parseEr := strconv.Atoi(args.Get("user_id"))
	day := args.Get("day")
	t, err := s.parseTime(day)
	if parseEr != nil || day == "" || err != nil {
		s.sendError(w, errors.New("bad input"))
	}
	ev, err := s.getByDay(userId, t)
	if err != nil {
		s.sendError(w, err)
	} else {
		json.NewEncoder(w).Encode(Result{ev})
	}
}

func (s *server) getByWeek(id int, day time.Time) ([]model.Event, error) {
	events, ok := s.storage.events[id]
	if !ok {
		return nil, errors.New("user doesnt exist")
	}
	res := make([]model.Event, 0, 10)
	for _, v := range events {
		w, y := v.Date.ISOWeek()
		needW, needY := day.ISOWeek()
		if w == needW && y == needY {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		return nil, errors.New("there are no events for this week")
	} else {
		return res, nil
	}
}

func (s *server) eventsWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	args := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")
	userId, parseEr := strconv.Atoi(args.Get("user_id"))
	day := args.Get("day")
	t, err := s.parseTime(day)
	if parseEr != nil || day == "" || err != nil {
		s.sendError(w, errors.New("bad input"))
	}
	ev, err := s.getByWeek(userId, t)
	if err != nil {
		s.sendError(w, err)
	} else {
		json.NewEncoder(w).Encode(Result{ev})
	}
}

func (s *server) getByMonth(id int, day time.Time) ([]model.Event, error) {
	events, ok := s.storage.events[id]
	if !ok {
		return nil, errors.New("user doesnt exist")
	}
	res := make([]model.Event, 0, 10)
	for _, v := range events {
		if (v.Date.Year() == day.Year()) && (v.Date.Month() == day.Month()) {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		return nil, errors.New("there are no events for this week")
	} else {
		return res, nil
	}
}

func (s *server) eventsMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	args := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")
	userId, parseEr := strconv.Atoi(args.Get("user_id"))
	day := args.Get("day")
	t, err := s.parseTime(day)

	if parseEr != nil || day == "" || err != nil {
		s.sendError(w, errors.New("bad input"))
		return
	}
	ev, err := s.getByMonth(userId, t)
	if err != nil {
		s.sendError(w, err)
	} else {
		json.NewEncoder(w).Encode(Result{ev})
	}
}
