package codes

var (
	// Basic
	ErrBadRequest = New(4000, "bad request")
	ErrAuthFailed = New(4001, "authentication failed")
	ErrNotFound   = New(4004, "not found")
	ErrDatabase   = New(5001, "database error")
	ErrInternal   = New(5000, "internal server error")

	// Common
	ErrInvalidUUID        = New(4001, "invalid uuid")
	ErrInvalidCredentials = New(4002, "invalid login or password")
	ErrJSONRequest        = New(4003, "invalid api request")
	ErrAlreadyExists      = New(4004, "aleady exsits")
	ErrConflict           = New(4005, "invalid request")

	// User
	ErrInvalidUserName = New(4101, "invalid user name")
	ErrInvalidEmail    = New(4102, "invalid email")
	ErrInvalidPassword = New(4103, "invalid password")

	// Monitor
	ErrInvalidMonitorName    = New(4201, "invalid monitor name")
	ErrInvalidMonitorURL     = New(4202, "invalid monitor url")
	ErrInvalidMonitorType    = New(4203, "invalid monitor type")
	ErrInvalidMonitorMethod  = New(4204, "invalid monitor method")
	ErrInvalidMonitorRequest = New(4205, "invalid monitor request")

	// Runner
	ErrInvalidRunnerName     = New(4301, "invalid runner name")
	ErrInvalidRunnerRegion   = New(4302, "invalid runner region")
	ErrInvalidRunnerInterval = New(4303, "invalid runner interval")

	// Controller / Request
	ErrSearchParameters = New(4901, "no search parameters")

	// Auth
	ErrTokenExpired = New(4401, "token expired")
	ErrTokenInvalid = New(4402, "invalid token")
)
