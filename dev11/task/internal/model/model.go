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
	Time time.Time `json:"date"`
}

func NewEvent(time time.Time) *Event {
	return &Event{
		Time: time,
	}
}

func NewUser() *User {
	defer atomic.AddInt64(&id, 1)
	return &User{ID: id}
}
