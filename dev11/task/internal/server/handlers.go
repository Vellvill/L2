package server

import (
	log "dev11/internal/logger "
	event "dev11/internal/model"
	"dev11/internal/usercases"
	"dev11/internal/validation"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Implementation struct {
	repo   usercases.Repository
	logger log.LoggerEx
}

func New(repo usercases.Repository, logger log.LoggerEx) Implementation {
	return Implementation{repo: repo, logger: logger}
}

func (i *Implementation) Create(w http.ResponseWriter, r *http.Request) {

	id, date, err := validation.ParseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	m, err := event.NewEvent(int64(id), date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model, err := i.repo.Create(m.ID, m.Date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status, err := w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

}

func (i *Implementation) Update(w http.ResponseWriter, r *http.Request) {

}

func (i *Implementation) Delete(w http.ResponseWriter, r *http.Request) {
	id, time, err := validation.ParseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = i.repo.Delete(int64(id), time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write([]byte("Event deleted"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (i *Implementation) Today(w http.ResponseWriter, r *http.Request) {

	id, _, err := validation.ParseParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := i.repo.Today(int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (i *Implementation) Middleware(next http.HandlerFunc, l log.LoggerEx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(r.URL.Query()["id"][0])
		if err != nil {
			i.logger.WriteErr(err)
		}

		if !validation.ValidateID(id) {
			http.Error(w, "err", http.StatusBadRequest)
		}

		reqInfo := fmt.Sprintf("Method: %s, id: %s", r.Method, r.URL.Query()["id"][0])
		err = l.WriteInfo(reqInfo)
		if err != nil {
			i.logger.WriteErr(err)
		}

		i.repo.Check(int64(id))

		next(w, r)
	}
}
