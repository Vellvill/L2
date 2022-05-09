package validation

import (
	"fmt"
	"net/http"
)

func ParseParams(r *http.Request, keys ...string) (map[string]string, error) {
	m := make(map[string]string)
	for _, key := range keys {
		keys, ok := r.URL.Query()[key]
		if !ok {
			return nil, fmt.Errorf("No %s in URL\n", keys)
		} else {
			m[key] = keys[0]
		}
	}
	return m, nil
}
