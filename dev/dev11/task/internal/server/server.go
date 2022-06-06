package server

import (
	cfg "dev11/internal/config"
	logaram "dev11/internal/logging"
	"dev11/internal/usercases"
	"fmt"
	"net"
	"net/http"
)

const (
	loggingPath = "./internal/logging/logs/logs.json"
)

type Application struct {
	mux    *http.ServeMux
	hash   usercases.Repository
	config *cfg.Cfg
	logger logaram.LoggerEx
}

func NewApp(config *cfg.Cfg, repo usercases.Repository) *Application {
	return &Application{
		hash:   repo,
		config: config,
	}
}

func (a *Application) Start() error {

	listener, err := net.Listen("tcp", net.JoinHostPort(a.config.Ip, a.config.Port))
	if err != nil {
		return err
	}

	a.logger, err = logaram.NewLogger(loggingPath)

	a.mux = NewServ(a.hash, a.logger)

	err = a.logger.WriteInfo(fmt.Sprintf("Starting server on %s:%s", a.config.Ip, a.config.Port))
	if err != nil {
		return err
	}

	return http.Serve(listener, a.mux)
}
