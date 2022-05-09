package hash

import (
	"dev11/internal/model"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Hash struct {
	sync.RWMutex
	id   int64
	hash map[int64]map[time.Time][]*model.Event
}

func NewHash() (*Hash, error) {
	return &Hash{
		RWMutex: sync.RWMutex{},
		id:      0,
		hash:    make(map[int64]map[time.Time][]*model.Event),
	}, nil
}

func (h *Hash) Check(id int64) (message string) {
	if _, ok := h.hash[id]; ok {
		return fmt.Sprintf("Id %d exists\n", id)
	} else {
		atomic.AddInt64(&h.id, 1)
		h.hash[id] = nil
		return fmt.Sprintf("Your new ID is %d\n", id)
	}
}

func (h *Hash) Create(id int64, date time.Time) (*model.Event, error) {
	if user, ok := h.hash[id]; ok {
		event := model.NewEvent(date)
		user[date] = append(user[date], event)
		return event, nil
	}
	return nil, fmt.Errorf("didn't find any users with id: %d\n", id)
}

func (h *Hash) Update(id int64, date time.Time, eventTime time.Time, newTime time.Time) (*model.Event, error) {
	if user, ok := h.hash[id]; ok {
		if j, ok := user[date]; ok {
			for _, v := range j {
				if v.Time == eventTime {
					v.Time = newTime
					return v, nil
				}
			}
		} else {
			return nil, fmt.Errorf("didn't find any events for %v\n", eventTime)
		}
	}
	return nil, fmt.Errorf("didn't find any users with id: %d\n", id)
}

func (h *Hash) Delete(id int64, date time.Time, eventTime time.Time) error {
	if user, ok := h.hash[id]; ok {
		if j, ok := user[date]; ok {
			for _, v := range j {
				if v.Time == eventTime {

					return nil
				}
			}
		} else {
			return fmt.Errorf("didn't find any events for %v\n", eventTime)
		}
	}
	return fmt.Errorf("didn't find any users with id: %d\n", id)
}

func (h *Hash) Today() ([]*model.Event, error) {
	return nil, nil
}

func (h *Hash) Week() ([]*model.Event, error) {
	return nil, nil
}

func (h *Hash) Month() ([]*model.Event, error) {
	return nil, nil
}
