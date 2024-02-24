package permission

import (
	"assignment-permission/cmd/server/models"
	"context"
)

type persistence interface {
	rolePersistence
}

type rolePersistence interface {
	InsertRole(ctx context.Context, prd *role) error
	InsertPermission(ctx context.Context, prd *permission) error
	getAllPermissions(ctx context.Context) (permissions, error)
}

type service struct {
	store persistence
}

type Service interface {
	InsertRole(ctx context.Context) error
	InsertPermission(ctx context.Context, p *permission) error
	GetAllPermissions(ctx context.Context) (models.Permissions, error)
}

func (svc *service) InsertRole(ctx context.Context) error {
	return nil
}

func (svc *service) InsertPermission(ctx context.Context, prm *permission) error {
	err := svc.store.InsertPermission(ctx, prm)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) GetAllPermissions(ctx context.Context) (models.Permissions, error) {
	prms, err := svc.store.getAllPermissions(ctx)
	if err != nil {
		return nil, err
	}
	domain, err := prms.Domain()
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func NewService(str persistence) (Service, error) {
	return &service{
		store: str,
	}, nil
}
