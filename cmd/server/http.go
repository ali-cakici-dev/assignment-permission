package server

import (
	"assignment-permission/cmd/server/models"
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
	fmt.Println("GetPermissions")
	permissions, err := ht.api.GetAllPermissions(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(permissions)
}

func (ht *HTTP) FetchPermittedResources(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	permissions, err := ht.api.FetchPermittedResources(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if permissions == nil {
		http.Error(w, "No permissions found for the user", http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(permissions)
}

func (ht *HTTP) GetRole(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	groupID := chi.URLParam(r, "group_id")
	permissions, err := ht.api.GetRole(r.Context(), userID, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if permissions == nil {
		http.Error(w, "No permissions found for the user", http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(permissions)
}

func (ht *HTTP) CreatePermission(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var newPermission models.Permission
	err := decoder.Decode(&newPermission)
	if err != nil {
		http.Error(w, "Error decoding permission data", http.StatusBadRequest)
		return
	}
	err = newPermission.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ht.api.InsertPermission(r.Context(), newPermission)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Permission created successfully"))
}

func (ht *HTTP) CreateRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var newRole models.Role
	err := decoder.Decode(&newRole)
	if err != nil {
		http.Error(w, "Error decoding permission data", http.StatusBadRequest)
		return
	}
	err = newRole.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ht.api.InsertRole(r.Context(), newRole)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Permission created successfully"))
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
	router.Get("/all-permissions", ht.GetPermissions)
	router.Post("/permissions", ht.CreatePermission)
	router.Post("/add-rold", ht.CreateRole)
	router.Get("/fetch-permitted-resources/{user_id}", ht.FetchPermittedResources)
	router.Get("/get-role/{group_id}/{user_id}", ht.GetRole)

	return ht
}
