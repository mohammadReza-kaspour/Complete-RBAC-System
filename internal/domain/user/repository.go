package user

import (
	"context"

	"github.com/mohammadReza-kaspour/Complete-RBAC-System/internal/domain/role"
)

type ListFilter struct {
	Search   string
	Page     int
	PageSize int
}

type Repository interface {
	// user
	Create(ctx context.Context, user *User) error
	FindById(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByMobile(ctx context.Context, mobile string) (*User, error)
	List(ctx context.Context, listFilter *ListFilter) ([]User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)

	// user roles
	AssignRole(ctx context.Context, userId string, roleId string) error
	RemoveRole(ctx context.Context, userId string, roleId string) error
	Roles(ctx context.Context, userId string) ([]role.Role, error)
	HasRole(ctx context.Context, userId string, roleName string) (bool, error)
}
