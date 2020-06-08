package model

type AuthUserData struct {
	Provider   string `json:"provider"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Identifier string `json:"identifier"`
}
