package server

import (
	"dev11/internal/usercases"
	"dev11/internal/validation"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Implementation struct {
	repo usercases.Repository
}

func New(repo usercases.Repository) Implementation {
	return Implementation{repo: repo}
}

type Req struct {
	Id   int       `json:"id"`
	Date time.Time `json:"date"`
}

func (r *Req) Read(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *Req) Close() error {
	//TODO implement me
	panic("implement me")
}

func newReq(id int, date time.Time) *Req {
	return &Req{Id: id, Date: date}
}

func (r *Req) validate() error {
	return nil
}

type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode int
	Req        *Req
}

func (i *Implementation) Create(w http.ResponseWriter, r *http.Request) {
	m, err := validation.ParseParams(r, "id", "date")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data, err := time.Parse("2006-01-02", m["date"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	id, err := strconv.Atoi(m["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	req := newReq(id, data)

	model, err := i.repo.Create(int64(req.Id), req.Date)
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

func (i *Implementation) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m, err := validation.ParseParams(r, "id", "date")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		id, err := strconv.Atoi(m["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		date, err := time.Parse("2006-01-02", m["date"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		req := newReq(id, date)

		message := i.repo.Check(int64(req.Id))

		log.Printf("%s\n", message)

		r.Body = req

		wi := &responseWriterInterceptor{
			statusCode: http.StatusOK,
			Req:        req,
		}

		next(wi, r)
	}
}
