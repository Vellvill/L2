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

func (h *Hash) Update(id int64, date time.Time, newTime time.Time) error {
	indexesForUpdate := make([]int, 0)
	h.RLock()
	insecure := h.hash[id]
	h.RUnlock()
	for i, v := range insecure {
		if v.Date == date {
			indexesForUpdate = append(indexesForUpdate, i)
		}
	}
	if len(indexesForUpdate) == 0 {
		return fmt.Errorf("didnt'")
	}

	h.update(indexesForUpdate, id, newTime)

	return nil
}

func (h *Hash) Delete(id int64, date time.Time) error {
	indexesForDel := make([]int, 0)
	h.RLock()
	insecure := h.hash[id]
	h.RUnlock()
	for i, v := range insecure {
		if v.Date == date {
			indexesForDel = append(indexesForDel, i)
		}
	}
	if len(indexesForDel) == 0 {
		return fmt.Errorf("no events with %v date", date)
	}

	h.delete(indexesForDel, id)

	return nil
}

func (h *Hash) Today(id int64) ([]byte, error) {
	return checkTimeDay(h.hash[id])
}

func (h *Hash) Week(id int64) ([]byte, error) {
	return checkTimeWeek(h.hash[id])
}

func (h *Hash) Month(id int64) ([]byte, error) {
	return checkTimeMonth(h.hash[id])
}
