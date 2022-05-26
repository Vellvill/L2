package hash

import (
	"dev11/internal/model"
	"time"
)

func (h *Hash) delete(indx int, id int64) error {
	h.hash[id][len(h.hash[id])-1] = h.hash[id][indx]
	h.hash[id] = h.hash[id][:len(h.hash)-1]
	return nil
}

func checkTime(date time.Time, which string, userEvents []*model.Event) []*model.Event {
	dayNow, monthNow, yearNow := time.Now().Day(), time.Now().Month(), time.Now().Year()
	time.Now().Weekday()
	result := make([]*model.Event, 0)
	switch which {
	case "today":
		for i, v := range userEvents {
			if v.Date.Year() == yearNow && v.Date.Month() == monthNow && v.Date.Day() == dayNow {
				result = append(result, userEvents[i])
			}
		}
	case "month":
		for i, v := range userEvents {
			if v.Date.Year() == yearNow && v.Date.Month() == monthNow {
				result = append(result, userEvents[i])
			}
		}
	case "week":
	default:
		return nil
	}
	return result
}
