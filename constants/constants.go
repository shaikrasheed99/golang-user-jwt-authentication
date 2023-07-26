package constants

const (
	Success            = "success"
	Error              = "error"
	Username           = "username"
	Role               = "role"
	Admin              = "admin"
	User               = "user"
	Authorization      = "Authorization"
	AccessTokenCookie  = "access_token"
	RefreshTokenCookie = "refresh_token"
	LocalHost          = "localhost"
	HomePath           = "/"
)

const (
	ErrUserAlreadyExists      = "user is already exists with username"
	ErrUserNotFound           = "user is not found with username"
	ErrUserIsNotAuthorised    = "user is not authorised to this api"
	ErrInvalidUsername        = "invalid username"
	ErrWrongPassword          = "password is wrong"
	ErrInvalidToken           = "invalid token"
	ErrExpiredToken           = "token has expired"
	ErrNoAuthHeader           = "no authorization header provided"
	ErrTokensNotFound         = "tokens are not found"
	ErrMaliciousToken         = "malicious token has been passed"
	ErrInvalidTokenExpiration = "invalid jwt access token expiration in minutes value"
)
