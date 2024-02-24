package api

import (
	"context"
)

func (ap *API) InsertPermission(ctx context.Context) error {
	err := ap.pService.InsertPermission(ctx)
	if err != nil {
		return err
	}

	return nil
}
