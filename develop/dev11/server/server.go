package server

import (
	"log"
	"net/http"
	"os"
	"server/middleware"
	"server/model"
	"server/repository"
)

type server struct {
	storage repository.Repository
	logger  *middleware.Logger
}

func newServer() *server {
	newServ := &server{}
	newServ.logger = &middleware.Logger{W: os.Stdout}
	newServ.storage = &repository.Cache{Events: make(map[int][]model.Event, 10)}
	return newServ
}

// Start функция старта http сервера
// парсит конфиг файл, запускает сервер в соответствии с конфигом, роутит запросы
func Start() {
	port := parseConfig()
	s := newServer()
	http.HandleFunc("/create_event", s.logger.Log(s.createEvent))
	http.HandleFunc("/update_event", s.logger.Log(s.updateEvent))
	http.HandleFunc("/delete_event", s.logger.Log(s.deleteEvent))
	http.HandleFunc("/events_for_day", s.logger.Log(s.eventsDay))
	http.HandleFunc("/events_for_week", s.logger.Log(s.eventsWeek))
	http.HandleFunc("/events_for_month", s.logger.Log(s.eventsMonth))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
