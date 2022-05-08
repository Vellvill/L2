package usercases

import (
	"dev11/internal/model"
	"time"
)

type Repository interface {
	Create(id int64, desc string, date time.Time) (*model.Event, error)
	Update(id int64, desc string, date time.Time, eventTime time.Time, newTime time.Time) (*model.Event, error)
	Delete(id int64, desc string, date time.Time, eventTime time.Time) error
	Today() ([]*model.Event, error)
	Week() ([]*model.Event, error)
	Month() ([]*model.Event, error)
}
