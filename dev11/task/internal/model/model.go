package model

import (
	"sync/atomic"
	"time"
)

var id int64

type User struct {
	ID int64 `json:"user_id"`
}

type Event struct {
	Time       time.Time `json:"date"`
	Description string    `json:"description"`
}

func NewEvent(time time.Time, desc string) *Event {
	return &Event{
		Time:        time,
		Description: desc,
	}
}

func NewUser() *User {
	defer atomic.AddInt64(&id, 1)
	return &User{ID: id}
}
