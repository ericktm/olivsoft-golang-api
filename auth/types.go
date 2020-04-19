package auth

// Credentials struct used to parse login requests
type Credentials struct {
	Username string `json:"username" binding:"required" `
	Password string `json:"password" binding:"required" `
}
