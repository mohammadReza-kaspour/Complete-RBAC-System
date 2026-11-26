package role

import (
	"context"

	"github.com/mohammadReza-kaspour/Complete-RBAC-System/internal/domain/permission"
)

type ListFilter struct {
	Search   string
	Page     int
	PageSize int
}

type Repository interface {
	// role
	Create(ctx context.Context, role *Role) error
	FindById(ctx context.Context, id string) (*Role, error)
	FindByName(ctx context.Context, name string) (*Role, error)
	List(ctx context.Context, listFilter ListFilter) (*Role, int64, error)
	Update(ctx context.Context, role *Role) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)

	// role - permission
	AssignPermission(ctx context.Context, roleId string, permissionId string) error
	RemovePermission(ctx context.Context, roleId string, permissionId string) error
	Permissions(ctx context.Context, roleId string) ([]permission.Permission, error)

	// user -> user_roles -> roles -> role_permissions -> permissions
	PermissionsOfUser(ctx context.Context, userId string) ([]permission.Permission, error)
	UserHasPermission(ctx context.Context, userId string, permissionId string) (bool, error)
}
