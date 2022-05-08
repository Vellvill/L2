package server

import (
	"dev11/internal/usercases"
	"net/http"
)

type Application struct {
	hash     usercases.Repository
	handlers http.Handler
}

func Start(port string, repo usercases.Repository) error {

	impl := New(repo)

	http.HandleFunc("/create", impl.Create)

	http.HandleFunc("/delete", impl.Delete)

	http.HandleFunc("/update", impl.Update)

	return http.ListenAndServe(port, nil)
}

func Middleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return handlerFunc
}
