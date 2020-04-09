package dto

type SessionInfoDto struct {
	AccountID	 string `json:"account_id,omitempty"`
	SessionToken string `json:"session_token,omitempty"`
}
