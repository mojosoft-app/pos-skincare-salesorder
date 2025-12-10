package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims represents JWT claims
type Claims struct {
	UserID      uint     `json:"user_id"`
	Name        string   `json:"name,omitempty"`
	RoleID      int      `json:"role_id,omitempty"`
	TenantCode  string   `json:"tenant_code"`
	IDLocation  *int     `json:"id_location,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	jwt.RegisteredClaims
}

// TokenResponse represents the JWT token response
type TokenResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int64     `json:"expires_in"`
	ExpiresAt    time.Time `json:"expires_at"`
}
