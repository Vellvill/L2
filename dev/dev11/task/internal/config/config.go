package config

import (
	"github.com/joho/godotenv"
	"os"
	"sync"
)

var (
	once sync.Once
	conf *Cfg
)

type Cfg struct {
	Port string
	Ip   string
}

func NewConfig() *Cfg {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
		conf = &Cfg{
			Port: os.Getenv("APP_PORT"),
			Ip:   os.Getenv("APP_HOST"),
		}

	})
	return conf
}
