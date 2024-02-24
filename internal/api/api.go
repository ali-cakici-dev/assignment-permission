package api

import (
	"assignment-permission/internal/permission"
)

type API struct {
	pService permission.Service
}

func New(pSvc permission.Service) *API {
	return &API{
		pService: pSvc,
	}
}
