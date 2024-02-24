package api

import (
	"context"
)

func (ap *API) InsertPermission(ctx context.Context) error {
	err := ap.pService.InsertRole(ctx)
	if err != nil {
		return err
	}

	return nil
}
