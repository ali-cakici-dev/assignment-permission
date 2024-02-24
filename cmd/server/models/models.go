package models

import "errors"

type Role struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Action      []string `json:"action"`
}

type Permissions []Permission
type Permission struct {
	Resource string           `json:"resource"`
	Users    UserPermissions  `json:"users"`
	Groups   GroupPermissions `json:"groups"`
}

func (p *Permission) Validate() error {
	if p.Resource == "" {
		return errors.New("resource is required")
	}
	if err := p.Users.validate(); err != nil {
		return err
	}
	if err := p.Groups.validate(); err != nil {
		return err
	}
	return nil
}

type UserID string
type UserIDs []UserID

type UserPermissions []UserPermission
type UserPermission struct {
	UserID UserID `json:"user_id"`
	Role   string `json:"role"`
}

func (u *UserPermissions) validate() error {
	if *u == nil {
		return errors.New("users are required")
	}
	for _, v := range *u {
		if v.UserID == "" {
			return errors.New("user id is required")
		}
		if v.Role == "" {
			return errors.New("role is required")
		}
	}
	return nil
}

type GroupPermissions []GroupPermission
type GroupPermission struct {
	GroupID string  `json:"group_id"`
	RoleID  string  `json:"role_id"`
	Members UserIDs `json:"member_ids"`
}

func (g *GroupPermissions) validate() error {
	if *g == nil {
		return errors.New("groups are required")
	}
	for _, v := range *g {
		if v.GroupID == "" {
			return errors.New("group id is required")
		}
		if v.RoleID == "" {
			return errors.New("role id is required")
		}
		if len(v.Members) == 0 {
			return errors.New("members are required")
		}
	}
	return nil
}
