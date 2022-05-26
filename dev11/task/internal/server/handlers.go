package server

import (
	event "dev11/internal/model"
	"dev11/internal/usercases"
	"dev11/internal/validation"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Implementation struct {
	repo usercases.Repository
}

func New(repo usercases.Repository) Implementation {
	return Implementation{repo: repo}
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

	var response struct {
		Models []event.Event `json:"events"`
	}

	for i := 0; i < len(res); i++ {
		response.Models = append(response.Models, *res[i])
	}

	resJson, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write(resJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (i *Implementation) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			fmt.Println(err)
		}

		i.repo.Check(int64(id))

		next(w, r)
	}
}
