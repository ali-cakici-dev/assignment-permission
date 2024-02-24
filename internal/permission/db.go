package permission

import (
	"assignment-permission/internal/pkg"
	"context"
)

type MongoConfig struct {
	PermissionCollection string
}

type PermissionDB struct {
	cli *pkg.MongoInstance
	cfg *MongoConfig
}

func (db *PermissionDB) InsertRole(ctx context.Context, prd *role) error {
	return nil
}

func InitDB(cli *pkg.MongoInstance, cfg *MongoConfig) (*PermissionDB, error) {

	return &PermissionDB{
		cli: cli,
		cfg: cfg,
	}, nil
}

func fetchDocument(res *document, collection string) (err error) {

	return
}
