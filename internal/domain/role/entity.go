package role

import (
	"regexp"
	"strings"
	"time"
	"uuid"

	"github.com/mohammadReza-kaspour/Complete-RBAC-System/internal/domain/domainerr"
	"github.com/mohammadReza-kaspour/Complete-RBAC-System/internal/domain/permission"
)

var (
	nameRe = regexp.MustCompile(`^[a-z][a-z0-9_]{1,63}$`)
)

type Role struct {
	Id          string
	Name        string
	Description string
	Permissions []permission.Permission
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func New(name string, description string) (*Role, error) {
	now := time.Now()
	role := &Role{
		Id:          uuid.NewV7().String(),
		Name:        name,
		Description: description,
		Permissions: []permission.Permission{},
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := role.Validate(); err != nil {
		return nil, err
	}
	return role, nil
}

func (r *Role) Validate() error {
	if !nameRe.MatchString(r.Name) {
		return domainerr.New(domainerr.InvalidRoleName, "invalid role name")
	}
	return nil
}
func (r *Role) Rename(name string) error {
	name = strings.ToLower(strings.TrimSpace(name))
	if !nameRe.MatchString(name) {
		return domainerr.New(domainerr.InvalidRoleName, "invalid role name")
	}
	r.Name = name
	return nil
}
func (r *Role) SetDescription(description string) {
	description = strings.TrimSpace(description)
	r.Description = description
}
func (r *Role) HasPermission(name string) bool {
	for _, permission := range r.Permissions {
		if permission.Name == name {
			return true
		}
	}
	return false
}
