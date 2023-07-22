package requests

type LogoutRequest struct {
	Username string `json:"username" binding:"required"`
}
