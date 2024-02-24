package permission

import (
	"assignment-permission/internal/pkg"
	"context"
	"go.mongodb.org/mongo-driver/bson"
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

func (db *PermissionDB) InsertPermission(ctx context.Context, prd *permission) error {
	permissionCollection := db.cli.DB.Collection(db.cfg.PermissionCollection)
	_, err := permissionCollection.InsertOne(ctx, prd)
	if err != nil {
		return err
	}
	return nil
}

func (db *PermissionDB) getAllPermissions(ctx context.Context) (permissions, error) {
	permissionCollection := db.cli.DB.Collection(db.cfg.PermissionCollection)
	find, err := permissionCollection.Find(
		ctx,
		bson.M{},
	)
	if err != nil {
		return nil, err
	}

	result := make([]permission, 0)
	for find.Next(ctx) {
		var p permission
		err := find.Decode(&p)
		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}

	return result, nil
}

func InitDB(cli *pkg.MongoInstance, cfg *MongoConfig) (*PermissionDB, error) {

	return &PermissionDB{
		cli: cli,
		cfg: cfg,
	}, nil
}

//func fetchDocument(res *document, collection string) (err error) {
//
//	return
//}
