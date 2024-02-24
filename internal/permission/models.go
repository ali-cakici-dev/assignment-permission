package permission

import (
	"assignment-permission/cmd/server/models"
	"math/rand"
)

type role struct {
	ID          string   `bson:"_id"`
	Name        string   `bson:"name"`
	Description string   `bson:"description"`
	Action      []string `bson:"action"`
}

type permissions []permission
type permission struct {
	Resource string           `bson:"resource"`
	Users    userPermissions  `bson:"users"`
	Groups   groupPermissions `bson:"groups"`
}

type userID string
type userIDs []userID

func (u *userID) generate() (id userID) {
	// generate user id randomly
	return userID(rune(rand.Int()))
}

type userPermissions []userPermission
type userPermission struct {
	UserID userID `bson:"user_id"`
	Role   string `bson:"role"`
}

type groupPermissions []groupPermission
type groupPermission struct {
	GroupID string  `bson:"group_id"`
	RoleID  string  `bson:"role_id"`
	Members userIDs `bson:"member_ids"`
}

/*
--------------------------------------------------------
DOMAIN CONVERSIONS MAY SEEM REDUNDANT, BUT THEY ARE NOT AS PROJECT GROWS
--------------------------------------------------------
*/

func (p *permission) Domain() (*models.Permission, error) {
	users, err := p.Users.Domain()
	if err != nil {
		return nil, err
	}

	groups, err := p.Groups.Domain()
	if err != nil {
		return nil, err
	}
	return &models.Permission{
		Resource: p.Resource,
		Users:    users,
		Groups:   groups,
	}, nil
}

func (ps *permissions) Domain() (models.Permissions, error) {
	var prms models.Permissions
	for _, v := range *ps {
		prm, err := v.Domain()
		if err != nil {
			return nil, err
		}
		prms = append(prms, *prm)
	}
	return prms, nil
}

func (u *userPermission) Domain() (models.UserPermission, error) {
	return models.UserPermission{
		UserID: models.UserID(u.UserID),
		Role:   u.Role,
	}, nil
}

func (u *userPermissions) Domain() (models.UserPermissions, error) {
	var users models.UserPermissions
	for _, v := range *u {
		user, err := v.Domain()
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (g *groupPermission) Domain() (models.GroupPermission, error) {
	return models.GroupPermission{
		GroupID: g.GroupID,
		RoleID:  g.RoleID,
		Members: g.Members.Domain(),
	}, nil
}

func (g *groupPermissions) Domain() (models.GroupPermissions, error) {
	var groups models.GroupPermissions
	for _, v := range *g {
		group, err := v.Domain()
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (u *userIDs) Domain() models.UserIDs {
	var users models.UserIDs
	for _, v := range *u {
		users = append(users, models.UserID(v))
	}
	return users
}

func (r *role) Domain() (*models.Role, error) {
	return &models.Role{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		Action:      r.Action,
	}, nil
}

func ToPermission(p *models.Permission) (*permission, error) {
	return &permission{
		Resource: p.Resource,
		Users:    ToUserPermissions(p.Users),
		Groups:   ToGroupPermissions(p.Groups),
	}, nil
}

func ToUserPermission(u models.UserPermission) userPermission {
	return userPermission{
		UserID: userID(u.UserID),
		Role:   u.Role,
	}
}

func ToUserPermissions(u models.UserPermissions) userPermissions {
	var users userPermissions
	for _, v := range u {
		users = append(users, ToUserPermission(v))
	}
	return users
}

func ToGroupPermission(g models.GroupPermission) groupPermission {
	return groupPermission{
		GroupID: g.GroupID,
		RoleID:  g.RoleID,
		Members: ToUserIDs(g.Members),
	}
}

func ToGroupPermissions(g models.GroupPermissions) groupPermissions {
	var groups groupPermissions
	for _, v := range g {
		groups = append(groups, ToGroupPermission(v))
	}
	return groups
}

func ToUserIDs(u models.UserIDs) userIDs {
	var users userIDs
	for _, v := range u {
		users = append(users, userID(v))
	}
	return users
}
