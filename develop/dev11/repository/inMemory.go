package repository

import (
	"errors"
	"server/model"
	"sync"
	"time"
)

// Cache стандартное inMemory хранилище
type Cache struct {
	sync.RWMutex
	Events map[int][]model.Event
}

func (c *Cache) addEvent(event model.Event) {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()
	c.Events[event.UserID] = append(c.Events[event.UserID], event)
}

func (c *Cache) updateEventStorage(event model.Event) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()
	if _, ok := c.Events[event.UserID]; !ok {
		return errors.New("user doesnt exist")
	}
	e := c.Events[event.UserID]
	for i := 0; i < len(e); i++ {
		if e[i].Date == event.Date && e[i].Title == event.Title {
			e[i] = event
			return nil
		}
	}

	return errors.New("event doesnt exist")
}

// Create создание записи
func (c *Cache) Create(request model.Request) error {
	e, err := validateReq(request)
	if err != nil {
		return err
	}
	c.addEvent(e)
	return nil
}

// Update обновление записи
func (c *Cache) Update(request model.Request) error {
	e, err := validateReq(request)
	if err != nil {
		return err
	}
	err = c.updateEventStorage(e)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache) deleteEventStorage(event model.Event) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()
	if _, ok := c.Events[event.UserID]; !ok {
		return errors.New("user doesnt exist")
	}
	e := c.Events[event.UserID]
	for i := 0; i < len(e); i++ {
		if e[i].Date == event.Date && e[i].Title == event.Title {
			e = append(e[:i], e[i+1:]...)
			c.Events[event.UserID] = e
			return nil
		}
	}

	return errors.New("event doesnt exist")
}

// Delete удаление записи
func (c *Cache) Delete(request model.Request) error {
	e, err := validateReq(request)
	if err != nil {
		return err
	}
	err = c.deleteEventStorage(e)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache) getByDay(id int, day time.Time) ([]model.Event, error) {
	events, ok := c.Events[id]
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
	}

	return res, nil
}

// GetByDay события на день
func (c *Cache) GetByDay(id int, date string) ([]model.Event, error) {
	t, err := parseTime(date)
	if err != nil {
		return nil, err
	}
	events, err := c.getByDay(id, t)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (c *Cache) getByWeek(id int, day time.Time) ([]model.Event, error) {
	events, ok := c.Events[id]
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
	}

	return res, nil
}

// GetByWeek события на неделю
func (c *Cache) GetByWeek(id int, date string) ([]model.Event, error) {
	t, err := parseTime(date)
	if err != nil {
		return nil, err
	}
	events, err := c.getByWeek(id, t)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (c *Cache) getByMonth(id int, day time.Time) ([]model.Event, error) {
	events, ok := c.Events[id]
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
	}

	return res, nil
}

// GetByMonth события на месяц
func (c *Cache) GetByMonth(id int, date string) ([]model.Event, error) {
	t, err := parseTime(date)
	if err != nil {
		return nil, err
	}
	events, err := c.getByMonth(id, t)
	if err != nil {
		return nil, err
	}

	return events, nil
}
