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
	getPermissionByUserID(ctx context.Context, userID string) (permissions, error)
	getRoleIDByUserIDGroupID(ctx context.Context, userID string, groupID string) (string, error)
	getRoleByID(ctx context.Context, roleID string) (*role, error)
}

type service struct {
	store persistence
}

type Service interface {
	InsertRole(ctx context.Context) error
	InsertPermission(ctx context.Context, p *permission) error
	GetAllPermissions(ctx context.Context) (models.Permissions, error)
	FetchPermittedResources(ctx context.Context, userID string) (models.Permissions, error)
	GetRole(ctx context.Context, userID string, groupID string) (*models.Role, error)
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

func (svc *service) FetchPermittedResources(ctx context.Context, userID string) (models.Permissions, error) {
	prms, err := svc.store.getPermissionByUserID(ctx, userID)
	if err != nil {
		return nil, err

	}
	domain, err := prms.Domain()
	if err != nil {
		return nil, err
	}
	return domain, nil
}

func (svc *service) GetRole(ctx context.Context, userID string, groupID string) (*models.Role, error) {
	roleID, err := svc.store.getRoleIDByUserIDGroupID(ctx, userID, groupID)
	if err != nil {
		return nil, err
	}

	r, err := svc.store.getRoleByID(ctx, roleID)
	if err != nil {
		return nil, err
	}

	domain, err := r.Domain()
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
