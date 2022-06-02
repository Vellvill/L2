package hash

import (
	"dev11/internal/model"
	"fmt"
	"sync"
	"time"
)

type Hash struct {
	sync.RWMutex
	hash map[int64][]*model.Event
}

func NewHash() (*Hash, error) {
	return &Hash{
		RWMutex: sync.RWMutex{},
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
		h.Lock()
		h.hash[id] = make([]*model.Event, 0)
		h.Unlock()
		return fmt.Sprintf("Your new ID is %d\n", id)
	}
}

func (h *Hash) Create(id int64, date time.Time) (model.Event, error) {
	h.RLock()
	if _, ok := h.hash[id]; ok {
		h.RUnlock()
		nModel, err := model.NewEvent(id, date)
		if err != nil {
			return model.Event{}, err
		}
		h.Lock()
		h.hash[id] = append(h.hash[id], &nModel)
		h.Unlock()
		return nModel, nil
	}
	h.RUnlock()
	return model.Event{
		ID:   0,
		Date: time.Time{},
	}, fmt.Errorf("didn't find any users with id: %d\n", id)
}

func (h *Hash) Update(id int64, date time.Time, newTime time.Time) (*model.Event, error) {
	h.RLock()

	for _, v := range h.hash[id] {
		if v.Date == date {
			v.Date = newTime
		}
	}

	h.RUnlock()
	return nil, fmt.Errorf("didn't find any events with %d ID\n", id)
}

func (h *Hash) Delete(id int64, date time.Time) error {
	indexesForDel := make([]int, 0)
	for i, v := range h.hash[id] {
		if v.Date == date {
			indexesForDel = append(indexesForDel, i)
		}
	}

	fmt.Errorf("df").Error()

	h.delete(indexesForDel, id)

	return fmt.Errorf("No events for %v, for %v\n", id, date)
}

func (h *Hash) Today(id int64) ([]byte, error) {
	return checkTime("today", h.hash[id])
}

func (h *Hash) Week(id int64) ([]byte, error) {
	return checkTime("week", h.hash[id])
}

func (h *Hash) Month(id int64) ([]byte, error) {
	return checkTime("month", h.hash[id])
}
