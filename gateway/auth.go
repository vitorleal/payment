package gateway

// Authenticated represents the authenticated user data
type LoggedUser struct {
	CompanyId int `json:"companyId"`
	AppId     int `json:"appId"`
	UserId    int `json:"userId"`
}

// IsNotValid check if the looged user has all mandatorty fields
func (logged *LoggedUser) IsNotValid() bool {
	if logged.CompanyId != 0 && logged.AppId != 0 && logged.UserId != 0 {
		return false
	}

	return true
}
