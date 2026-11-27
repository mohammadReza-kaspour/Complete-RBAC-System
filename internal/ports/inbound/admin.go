package inbound

import "context"

type Overview struct {
	Users          int64 `json:"users"`
	Roles          int64 `json:"roles"`
	Permissions    int64 `json:"permissions"`
	ActiveSessions int64 `json:"activeSessions"`
}

type SeedResult struct {
	RolesCreated        int64 `json:"rolesCreated"`
	PermissionsCreated  int64 `json:"permissionsCreated"`
	PermissionsAssigned int64 `json:"permissionsAssigned"`
	AdminUserEnsured    bool  `json:"adminUserEnsured"`
}

type CheckPermissionInput struct {
	UserID     string `json:"userId"`
	Permission string `json:"permission"`
}

type AdminService interface {
	Overview(ctx context.Context) (*Overview, error)
	SeedDefaults(ctx context.Context) (*SeedResult, error)
	EnsureAdminUser(ctx context.Context, mobile string) (bool, error)
}
