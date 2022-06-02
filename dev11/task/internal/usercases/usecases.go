package usercases

import (
	"dev11/internal/model"
	"time"
)

type Repository interface {
	Check(id int64) (message string)
	Create(id int64, date time.Time) (model.Event, error)
	Update(id int64, date time.Time, newTime time.Time) (*model.Event, error)
	Delete(id int64, date time.Time) error
	Today(id int64) ([]byte, error)
	Week(id int64) ([]byte, error)
	Month(id int64) ([]byte, error)
}
