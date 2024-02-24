package service

import "context"

type document interface {
	fetch(context.Context, string) error
}

type role struct {
	ID          string   `bson:"_id"`
	Name        string   `bson:"name"`
	Description string   `bson:"description"`
	Action      []string `bson:"action"`
}

func (r *role) fetch(ctx context.Context, id string) (err error) {

	return
}

type permission struct {
	Resource string            `bson:"resource"`
	Users    []userPermission  `bson:"users"`
	Groups   []groupPermission `bson:"groups"`
}

func (r *permission) fetch(ctx context.Context, id string) (err error) {

	return
}

type userPermission struct {
	UserID string `bson:"user_id"`
	Role   string `bson:"role"`
}

type groupPermission struct {
	GroupID string   `bson:"group_id"`
	Role    string   `bson:"role"`
	Members []string `bson:"member_ids"`
}

func fetchPermittedResources(ctx context.Context, userID string) (resources []string, err error) {

	return
}

func checkPermission(ctx context.Context, userID, resourceID, action string) (err error) {

	return
}

func getRole(ctx context.Context, userID, resourceID string) (role string, err error) {

	return
}
