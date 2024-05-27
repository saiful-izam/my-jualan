package dto

type ErrorDto struct {
	Error string `json:"error"`
}

type SessionDTO struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
