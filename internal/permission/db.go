package permission

import (
	"assignment-permission/internal/pkg"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoConfig struct {
	PermissionCollection string
	RoleCollection       string
}

type PermissionDB struct {
	cli *pkg.MongoInstance
	cfg *MongoConfig
}

func (db *PermissionDB) insertPermission(ctx context.Context, prd *permission) error {
	permissionCollection := db.cli.DB.Collection(db.cfg.PermissionCollection)
	_, err := permissionCollection.InsertOne(ctx, prd)
	if err != nil {
		return err
	}
	return nil
}

func (db *PermissionDB) insertRole(ctx context.Context, r *role) error {
	permissionCollection := db.cli.DB.Collection(db.cfg.RoleCollection)
	_, err := permissionCollection.InsertOne(ctx, r)
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

func (db *PermissionDB) getPermissionByUserID(ctx context.Context, userID string) (permissions, error) {
	permissionCollection := db.cli.DB.Collection(db.cfg.PermissionCollection)
	find, err := permissionCollection.Find(
		ctx,
		bson.M{"users.user_id": userID},
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

func (db *PermissionDB) getRoleIDByUserIDGroupID(ctx context.Context, userID string, groupID string) (string, error) {
	permissionsCollection := db.cli.DB.Collection(db.cfg.PermissionCollection)
	res := permissionsCollection.FindOne(
		ctx,
		bson.M{
			"$and": bson.A{
				bson.M{userID: bson.M{"$in": "group.members"}},
				bson.M{"groups.group_id": groupID},
			},
		},
	)
	if res.Err() != nil {
		return "", res.Err()
	}

	var p permission
	err := res.Decode(&p)
	if err != nil {
		return "", err
	}

	for _, v := range p.Groups {
		if v.GroupID == groupID {
			return v.RoleID, nil
		}
	}

	return "", nil
}

func (db *PermissionDB) getRoleByID(ctx context.Context, roleID string) (*role, error) {
	rolesCollection := db.cli.DB.Collection(db.cfg.RoleCollection)
	res := rolesCollection.FindOne(
		ctx,
		bson.M{"_id": roleID},
	)
	if res.Err() != nil {
		return nil, res.Err()
	}

	var r role
	err := res.Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func InitDB(cli *pkg.MongoInstance, cfg *MongoConfig) (*PermissionDB, error) {

	return &PermissionDB{
		cli: cli,
		cfg: cfg,
	}, nil
}
