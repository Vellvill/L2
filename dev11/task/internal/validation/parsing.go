package validation

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func ParseParams(r *http.Request) (int, time.Time, error) {
	id, ok := r.URL.Query()["id"]
	if ok {
		date, ok := r.URL.Query()["date"]
		if ok {
			return validateParams(id[0], date[0])
		}
		return 0, time.Time{}, fmt.Errorf("bad request")
	}
	return 0, time.Time{}, fmt.Errorf("bad request")
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

	data, err = time.Parse("2006-01-02", date)
	if err != nil {
		return 0, time.Time{}, err
	}

	return idd, data, nil
}
