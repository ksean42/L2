package repository

import (
	"errors"
	"server/model"
	"time"
)

func parseTime(date string) (time.Time, error) {
	res, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, err
	}
	return res, nil
}

func validateReq(decoded model.Request) (model.Event, error) {

	t, errTime := parseTime(decoded.Date)
	if errTime != nil {
		return model.Event{}, errTime
	}
	if decoded.Title == "" || decoded.UserID == 0 || decoded.Description == "" {
		return model.Event{}, errors.New("bad request")
	}
	newEvent := model.Event{
		UserID:      decoded.UserID,
		Date:        t,
		Title:       decoded.Title,
		Description: decoded.Description}
	return newEvent, nil

}
