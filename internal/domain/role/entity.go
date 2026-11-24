package role

import (
	"time"

	"github.com/mohammadReza-kaspour/Complete-RBAC-System/internal/domain/permission"
)

type Role struct {
	Id          string
	Name        string
	Description string
	Permissions []permission.Permission
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
