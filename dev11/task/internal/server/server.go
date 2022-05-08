package server

import (
	"dev11/internal/usercases"
	"net"
	"net/http"
)

type Application struct {
	hash     usercases.Repository
	handlers http.Handler
}

func Start(port string, repo usercases.Repository) error {

	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	impl := New(repo)

	http.HandleFunc("/create", impl.Create)

	http.HandleFunc("/delete", impl.Delete)

	http.HandleFunc("/update", impl.Update)

	return http.Serve(listener, nil)
}

func Middleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return handlerFunc
}
