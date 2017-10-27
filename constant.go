package auth_backend

import (
	"crypto/rand"
	"encoding/base64"
)

const (
	// URL_AUTHORIZATION used to request an authorization
	URL_AUTHORIZATION = "/oauth2/authorize"
	// URL_ACCESS_TOKEN used to get a new access token
	URL_ACCESS_TOKEN = "/oauth2/access_token"
	// URL_USER_INFO used to retrieve current user information
	URL_USER_INFO = "/oauth2/user_info/"
	// URL_REDIRECT used as parameters when requesting an authorization as redirect uri
	URL_REDIRECT = "/complete/edx-oidc/"
)

// Scopes return a default scopes for authentication
func Scopes() []string {
	return []string{"profile", "email", "openid"}
}

// RandomToken generate a random string that will be used as state
func RandomToken(l int) string {
	b := make([]byte, l)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
