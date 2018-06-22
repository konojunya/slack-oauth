package service

type CredentialInfo struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	UserID      string `json:"user_id"`
	TeamName    string `json:"team_name"`
	TeamID      string `json:"team_id"`
}

var (
	Credential *CredentialInfo
)

func SetConfig(credential *CredentialInfo) {
	Credential = credential
}
