package outbound

import (
	"github.com/mohammadReza-kaspour/Complete-RBAC-System/internal/domain/otp"
	"github.com/mohammadReza-kaspour/Complete-RBAC-System/internal/domain/permission"
	"github.com/mohammadReza-kaspour/Complete-RBAC-System/internal/domain/role"
	"github.com/mohammadReza-kaspour/Complete-RBAC-System/internal/domain/session"
	"github.com/mohammadReza-kaspour/Complete-RBAC-System/internal/domain/user"
)

type Repositories struct {
	Users       user.Repository
	Roles       role.Repository
	Permissions permission.Repository
	OTPs        otp.Repository
	Sessions    session.Repository
}
