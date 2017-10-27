package auth_backend

// UserInfo stored user information from user endpoint edx
type UserInfo struct {
	TrackingID uint   `json:"user_tracking_id"`
	Email      string `json:"email"`
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	Locale     string `json:"locale"`
	Name       string `json:"name'`
	Username   string `json:"preferred_username"`
	Sub        string `json:"sub"`
}
