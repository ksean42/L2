package repository

import (
	"server/model"
)

// Repository интерфейс хранилища
type Repository interface {
	Create(request model.Request) error
	Update(request model.Request) error
	Delete(request model.Request) error
	GetByDay(id int, date string) ([]model.Event, error)
	GetByWeek(id int, date string) ([]model.Event, error)
	GetByMonth(id int, date string) ([]model.Event, error)
}
