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
	h.RLock()
	if _, ok := h.hash[id]; ok {
		h.RUnlock()
		return fmt.Sprintf("Id %d exists\n", id)
	} else {
		h.RUnlock()
		atomic.AddInt64(&h.id, 1)
		h.Lock()
		h.hash[id] = make(map[time.Time][]*model.Event)
		h.Unlock()
		return fmt.Sprintf("Your new ID is %d\n", id)
	}
}

func (h *Hash) Create(id int64, date time.Time) (*model.Event, error) {
	h.RLock()
	if user, ok := h.hash[id]; ok {
		h.RUnlock()
		event, err := model.NewEvent(id, date)
		if err != nil {
			return nil, err
		}
		h.Lock()
		user[date] = append(user[date], event)
		h.Unlock()
		return event, nil
	}
	h.RUnlock()
	return nil, fmt.Errorf("didn't find any users with id: %d\n", id)
}

func (h *Hash) Update(id int64, date time.Time, eventTime time.Time, newTime time.Time) (*model.Event, error) {
	if user, ok := h.hash[id]; ok {
		if j, ok := user[date]; ok {
			for _, v := range j {
				if v.Date == eventTime {
					v.Date = newTime
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
				if v.Date == eventTime {

					return nil
				}
			}
		} else {
			return fmt.Errorf("didn't find any events for %v\n", eventTime)
		}
	}
	return fmt.Errorf("didn't find any users with id: %d\n", id)
}

func (h *Hash) Today(id int64) ([]*model.Event, error) {
	if user, ok := h.hash[id]; ok {
		if j, ok := user[time.Now()]; ok {
			return j, nil
		} else {
			return nil, fmt.Errorf("didn't find any events for %v\n", time.Now)
		}
	}
	return nil, fmt.Errorf("didn't find any users with id: %d\n", id)
}

func (h *Hash) Week() ([]*model.Event, error) {
	return nil, nil
}

func (h *Hash) Month() ([]*model.Event, error) {
	return nil, nil
}
