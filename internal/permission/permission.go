package permission

import (
	"context"
	"math/rand"
)

type document interface {
	fetch(context.Context, string) error
}

type role struct {
	ID          string   `bson:"_id"` // mongodb ObjectID
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

type userID string
type userIDs []userID

func (u *userID) generate() (id userID) {
	// generate user id randomly
	return userID(rune(rand.Int()))
}

type userPermission struct {
	UserID userID `bson:"user_id"`
	Role   string `bson:"role"`
}

type groupPermission struct {
	GroupID string  `bson:"group_id"`
	RoleID  string  `bson:"role_id"`
	Members userIDs `bson:"member_ids"`
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
