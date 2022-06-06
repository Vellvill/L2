package hash

import (
	"dev11/internal/model"
	"encoding/json"
	"time"
)

var (
	dayNow, monthNow, yearNow = time.Now().Day(), time.Now().Month(), time.Now().Year()
	_, weekNow                = time.Now().ISOWeek()
)

func (h *Hash) delete(indx []int, id int64) {
	h.Lock()
	for _, i := range indx {
		h.hash[id][len(h.hash[id])-1] = h.hash[id][i]
		h.hash[id] = h.hash[id][:len(h.hash)-1]
	}
	h.Unlock()
}

func (h *Hash) update(indx []int, id int64, newTime time.Time) {
	h.Lock()
	for _, i := range indx {
		h.hash[id][i].Date = newTime
	}
	h.Unlock()
}

func checkTimeDay(userEvents []*model.Event) ([]byte, error) {
	result := make([]*model.Event, 0)
	for _, v := range userEvents {
		if v.Date.Day() == dayNow && v.Date.Month() == monthNow && v.Date.Year() == yearNow {
			result = append(result, v)
		}
	}
	return newJson(result)
}

func checkTimeWeek(userEvents []*model.Event) ([]byte, error) {
	result := make([]*model.Event, 0)
	for _, v := range userEvents {
		eventYear, eventWeek := v.Date.ISOWeek()
		if eventYear == yearNow && eventWeek == weekNow {
			result = append(result, v)
		}
	}
	return newJson(result)
}
func checkTimeMonth(userEvents []*model.Event) ([]byte, error) {
	result := make([]*model.Event, 0)
	for _, v := range userEvents {
		if v.Date.Month() == monthNow && v.Date.Year() == yearNow {
			result = append(result, v)
		}
	}
	return newJson(result)
}

type jsonEvents struct {
	Models []*model.Event `json:"events"`
}

func newJson(events []*model.Event) ([]byte, error) {
	var jsonevents jsonEvents

	for _, v := range events {
		jsonevents.Models = append(jsonevents.Models, v)
	}

	js, err := json.Marshal(jsonevents)
	if err != nil {
		return nil, err
	}

	return js, nil
}
