package permission

import (
	"context"
)

type persistence interface {
	rolePersistence
}

type rolePersistence interface {
	InsertRole(ctx context.Context, prd *role) error
}

type service struct {
	store persistence
}

type Service interface {
	InsertRole(ctx context.Context) error
}

func (svc *service) InsertRole(ctx context.Context) error {
	return nil
}

func NewService(str persistence) (Service, error) {
	return &service{
		store: str,
	}, nil
}
