package server

import (
	cfg "dev11/internal/config"
	logaram "dev11/internal/logger "
	"dev11/internal/usercases"
	"net"
	"net/http"
)

const (
	loggingPath = "./logs/logs.json"
)

type Application struct {
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

	impl := New(a.hash, a.logger)

	http.HandleFunc("/create", impl.Middleware(impl.Create, a.logger))

	http.HandleFunc("/delete", impl.Middleware(impl.Delete, a.logger))

	http.HandleFunc("/update", impl.Middleware(impl.Update, a.logger))

	http.HandleFunc("/today", impl.Middleware(impl.Today, a.logger))

	defer func() {
		<-a.logger.SaveWriting()
	}()

	return http.Serve(listener, nil)
}
