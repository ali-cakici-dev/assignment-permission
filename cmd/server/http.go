package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"net/http"
	"sync"
)

type Config struct {
	Host string
	Port int
}

type HTTP struct {
	lock   *sync.Mutex
	server *http.Server
}

func (ht *HTTP) Start() error {
	err := ht.server.ListenAndServe()
	if err != nil {
		return errors.Wrap(err, "http listen and serve ended")
	}
	return nil
}

func New(cfg *Config) *HTTP {
	ht := &HTTP{
		lock: &sync.Mutex{},
	}

	router := chi.NewRouter()
	ht.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler: router,
	}

	return ht
}
