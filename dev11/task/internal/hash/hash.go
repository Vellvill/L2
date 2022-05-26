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
	hash map[int64][]*model.Event
}

func NewHash() (*Hash, error) {
	return &Hash{
		RWMutex: sync.RWMutex{},
		id:      0,
		hash:    make(map[int64][]*model.Event, 0),
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
		h.hash[h.id] = make([]*model.Event, 0)
		h.Unlock()
		return fmt.Sprintf("Your new ID is %d\n", id)
	}
}

func (h *Hash) Create(id int64, date time.Time) (*model.Event, error) {
	h.RLock()
	if _, ok := h.hash[id]; ok {
		h.RUnlock()
		nModel, err := model.NewEvent(id, date)
		if err != nil {
			return nil, err
		}
		h.Lock()
		h.hash[id] = append(h.hash[id], nModel)
		h.Unlock()
		return nModel, nil
	}
	h.RUnlock()
	return nil, fmt.Errorf("didn't find any users with id: %d\n", id)
}

func (h *Hash) Update(id int64, date time.Time, newTime time.Time) (*model.Event, error) {
	for _, v := range h.hash[id] {
		if v.Date == date {
			v.Date = newTime
			return v, nil
		}
	}
	return nil, fmt.Errorf("didn't find any events with %d ID\n", id)
}

func (h *Hash) Delete(id int64, date time.Time) error {
	for i, v := range h.hash[id] {
		if v.Date == date {
			return h.delete(i, id)
		}
	}
	return fmt.Errorf("No events for %v, for %v\n", id, date)
}

func (h *Hash) Today(id int64) ([]*model.Event, error) {
	return checkTime(time.Now(), "today", h.hash[id]), nil
}

func (h *Hash) Week(id int64) ([]*model.Event, error) {
	return nil, nil
}

func (h *Hash) Month(id int64) ([]*model.Event, error) {
	return checkTime(time.Now(), "month", h.hash[id]), nil
}
