package server

import (
	"dev11/internal/usercases"
	"encoding/json"
	"net/http"
)

type Implementation struct {
	repo usercases.Repository
}

type Req struct {
	id          int64 `json:"id"`
	Date        `json:"date"`
	Description string `json:"description"`
}

func New(repo usercases.Repository) Implementation {
	return Implementation{repo: repo}
}

func (i *Implementation) Create(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	model, err := i.repo.Create(e.id, e.Description, e.Date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	res, err := json.Marshal(model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	status, err := w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), status)
	}
}

func (i *Implementation) Update(w http.ResponseWriter, r *http.Request) {

}

func (i *Implementation) Delete(w http.ResponseWriter, r *http.Request) {

}
