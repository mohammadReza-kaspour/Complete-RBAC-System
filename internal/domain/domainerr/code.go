package domainerr

const (
	// Auth context 1100-1199.
	InvalidToken       = 1100
	OTPInvalidCode     = 1102
	OTPExpired         = 1103
	OTPAlreadyUsed     = 1104
	OTPTooManyAttempts = 1105
	OTPNotFound        = 1106
	SessionExpired     = 1107
	SessionRevoked     = 1108
	AccountDisabled    = 1109
	SessionNotFound    = 1111

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
