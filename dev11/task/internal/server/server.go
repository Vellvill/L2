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

	http.HandleFunc("/create", impl.Middleware(impl.Create))

	http.HandleFunc("/delete", impl.Middleware(impl.Delete))

	http.HandleFunc("/update", impl.Middleware(impl.Update))

	return http.Serve(listener, nil)
}
