package inbound

import "context"

type CreateRoleInput struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	PermissionIds []string `json:"permissionIds"`
}
type UpdateRoleInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type CreatePermissionInput struct {
	Resource    string `json:"resource"`
	Action      string `json:"action"`
	Description string `json:"description"`
}
type UpdatePermissionInput struct {
	Resource    *string `json:"resource"`
	Action      *string `json:"action"`
	Description *string `json:"description"`
}

type RoleManagement interface {
	// Role
	Create(ctx context.Context, cmd CreateRoleInput) (*RoleView, error)
	Get(ctx context.Context, id string) (*RoleView, error)
	List(ctx context.Context, listQuery *ListQuery) (*Page[RoleView], error)
	Update(ctx context.Context, id string, cmd UpdateRoleInput) (*RoleView, error)
	Delete(ctx context.Context, id string) error

	// Role - Permission
	AssignPermissionToRole(ctx context.Context, roleId string, permissionId string) error
	RemovePermissionFromRole(ctx context.Context, roleId string, permissionId string) error
	RolePermissions(ctx context.Context, roleId string) ([]PermissionView, error)

	// Role - User
	AssignRoleToUser(ctx context.Context, userId string, roleId string) error
	RemoveRoleFromUser(ctx context.Context, userId string, roleId string) error
	UserRoles(ctx context.Context, userId string) ([]RoleView, error)
}
type PermissionManagement interface {
	// Permission
	Create(ctx context.Context, cmd CreatePermissionInput) (*PermissionView, error)
	Get(ctx context.Context, id string) (*PermissionView, error)
	List(ctx context.Context, listQuery *ListQuery) (*Page[PermissionView], error)
	Update(ctx context.Context, id string, cmd UpdatePermissionInput) (*PermissionView, error)
	Delete(ctx context.Context, id string) error

	// User - Permission
	UserPermissions(ctx context.Context, userID string) ([]PermissionView, error)
}

// for middleware
type PermissionChecker interface {
	UserHasPermission(ctx context.Context, userID uint, permission string) (bool, error)
}
