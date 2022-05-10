package model

import (
	"fmt"
	"time"
)

type Event struct {
	ID   int64     `json:"id"`
	Date time.Time `json:"date"`
}

func NewEvent(id int64, time time.Time) (*Event, error) {
	e := &Event{
		ID:   id,
		Date: time,
	}

	return e, e.validate()
}

func (e *Event) validate() error {
	if e.Date.Before(time.Now()) {
		return fmt.Errorf("Creating date before now\n")
	} else {
		return nil
	}
}
