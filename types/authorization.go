package types

type ParseTokenWorkstationResponse struct {
	UserId        int
	WorkstationId int
}

type SessionInfo struct {
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
	ExpiresAt    int    `json:"expires_at" db:"expires_at"`
}
