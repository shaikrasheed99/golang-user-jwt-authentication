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
	UserAlreadyExistsErrorMessage      = "user is already exists with username"
	UserNotFoundErrorMessage           = "user is not found with username"
	UserIsNotAuthorisedErrorMessage    = "user is not authorised to this api"
	InvalidUsernameErrorMessage        = "invalid username"
	WrongPasswordErrorMessage          = "password is wrong"
	InvalidTokenErrorMessage           = "invalid token"
	ExpiredTokenErrorMessage           = "token has expired"
	NoAuthHeaderErrorMessage           = "no authorization header provided"
	TokensNotFoundErrorMessage         = "tokens are not found"
	MaliciousTokenErrorMessage         = "malicious token has been passed"
	InvalidTokenExpirationErrorMessage = "invalid jwt access token expiration in minutes value"
)
