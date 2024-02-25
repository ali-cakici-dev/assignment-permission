package api

import (
	"assignment-permission/cmd/server/models"
	"assignment-permission/internal/permission"
	"context"
	"fmt"
)

func (ap *API) InsertPermission(ctx context.Context, p models.Permission) error {
	prm, err := permission.ToPermission(&p)
	if err != nil {
		return err
	}

	err = ap.pService.InsertPermission(ctx, prm)
	if err != nil {
		return err
	}

	return nil
}

func (ap *API) GetAllPermissions(ctx context.Context) (models.Permissions, error) {
	ps, err := ap.pService.GetAllPermissions(ctx)
	fmt.Println("ps: ", ps)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (ap *API) FetchPermittedResources(ctx context.Context, userID string) (models.Permissions, error) {
	ps, err := ap.pService.FetchPermittedResources(ctx, userID)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (ap *API) GetRole(ctx context.Context, userID string, groupID string) (*models.Role, error) {
	ps, err := ap.pService.GetRole(ctx, userID, groupID)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (ap *API) InsertRole(ctx context.Context, role models.Role) error {
	err := ap.pService.InsertRole(ctx)
	if err != nil {
		return err
	}

	return nil
}
