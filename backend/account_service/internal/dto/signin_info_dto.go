package dto

type SigninInfoDto struct {
	Email 	 string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
