package validation

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	br          = "bad request"
	timeExample = "2006-01-02T15:04:05"
)

func ParseParams(r *http.Request) (int, time.Time, error) {
	id, ok := r.URL.Query()["id"]
	if ok {
		date, ok := r.URL.Query()["date"]
		if ok {
			return validateParams(id[0], date[0])
		} else {
			idRes, err := strconv.Atoi(id[0])
			if err != nil {
				return 0, time.Time{}, err
			}
			return idRes, time.Time{}, nil
		}
	}
	return 0, time.Time{}, fmt.Errorf(br)
}

func validateParams(id string, date string) (int, time.Time, error) {
	var (
		idd  int
		data time.Time
	)

	idd, err := strconv.Atoi(id)
	if err != nil {
		return 0, time.Time{}, err
	}

	data, err = time.Parse(timeExample, date)
	if err != nil {
		return 0, time.Time{}, err
	}

	return idd, data, nil
}

func ValidateID(id int) bool {
	if id <= 0 {
		return false
	}
	return true
}
