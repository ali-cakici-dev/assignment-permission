package server

import (
	"assignment-permission/internal/api"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
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
	api    *api.API
}

func (ht *HTTP) Start() error {
	err := ht.server.ListenAndServe()
	if err != nil {
		return errors.New("failed to start http server")
	}
	return nil
}

func (ht *HTTP) GetPermissions(w http.ResponseWriter, r *http.Request) {
	//permissions, err := ht.api.GetPermissions()
	permissions := []string{"read", "write", "delete"}
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	_ = json.NewEncoder(w).Encode(permissions)
}

func New(cfg *Config, api *api.API) *HTTP {
	ht := &HTTP{
		lock: &sync.Mutex{},
		api:  api,
	}

	router := chi.NewRouter()
	ht.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler: router,
	}
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})
	router.Get("/permissions", ht.GetPermissions)

	return ht
}
