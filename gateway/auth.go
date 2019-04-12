package gateway

// Authenticated represents the authenticated user data
type LoggedUser struct {
	UserId int `json:"userId"`
}

// IsNotValid check if the looged user has all mandatorty fields
func (logged *LoggedUser) IsNotValid() bool {
	if logged.UserId != 0 {
		return false
	}

	return true
}
