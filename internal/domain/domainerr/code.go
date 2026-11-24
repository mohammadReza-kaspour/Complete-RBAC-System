package domainerr

const (
	// Auth context 1100-1199.
	InvalidToken = 1100

	// User context 1200-1299.
	UserNotFound      = 1200
	UserAlreadyExists = 1201
	InvalidUsername   = 1202
	InvalidEmail      = 1203
	InvalidMobile     = 1204

	// Role context 1300-1399.
	RoleNotFound      = 1300
	RoleAlreadyExists = 1301
	InvalidRoleName   = 1302

	// Permission context 1400-1499.
	PermissionNotFound      = 1400
	PermissionAlreadyExists = 1401
	InvalidPermissionParts  = 1402
)
