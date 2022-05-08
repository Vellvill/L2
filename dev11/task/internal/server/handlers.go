package server

import (
	"dev11/internal/usercases"
	"encoding/json"
	"net/http"
	"time"
)

type Implementation struct {
	repo usercases.Repository
}

type Event struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

func New(repo usercases.Repository) Implementation {
	return Implementation{repo: repo}
}

func (i *Implementation) Create(w http.ResponseWriter, r *http.Request) {
	var e Event
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (i *Implementation) Update(w http.ResponseWriter, r *http.Request) {

}

func (i *Implementation) Delete(w http.ResponseWriter, r *http.Request) {

}
