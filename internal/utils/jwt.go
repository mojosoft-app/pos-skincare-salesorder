package utils

import (
	"fmt"
	"time"

	"pos-mojosoft-so-service/internal/config"
	"pos-mojosoft-so-service/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

type JWTUtil struct {
	config *config.JWTConfig
}

func NewJWTUtil(config *config.JWTConfig) *JWTUtil {
	return &JWTUtil{
		config: config,
	}
}

func (j *JWTUtil) GenerateAccessToken(userID uint, name string, roleID int, tenantCode string, permissions []string, idLocation *int) (string, time.Time, error) {
	expiresAt := time.Now().Add(j.config.AccessTokenTTL)

	claims := models.Claims{
		UserID:      userID,
		Name:        name,
		RoleID:      roleID,
		TenantCode:  tenantCode,
		IDLocation:  idLocation,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    j.config.Issuer,
			Subject:   fmt.Sprintf("%d", userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.config.Secret))
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to generate access token: %w", err)
	}

	return tokenString, expiresAt, nil
}

func (j *JWTUtil) GenerateRefreshToken(userID uint, tenantCode string) (string, error) {
	expiresAt := time.Now().Add(j.config.RefreshTokenTTL)

	// Use custom claims to include tenant_code
	claims := models.Claims{
		UserID:     userID,
		TenantCode: tenantCode,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    j.config.Issuer,
			Subject:   fmt.Sprintf("%d", userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.config.Secret))
	if err != nil {
		return "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return tokenString, nil
}

func (j *JWTUtil) ValidateAccessToken(tokenString string) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.config.Secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}

func (j *JWTUtil) ValidateRefreshToken(tokenString string) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.config.Secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse refresh token: %w", err)
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid refresh token")
}

func (j *JWTUtil) CreateTokenResponse(accessToken, refreshToken string, expiresAt time.Time) models.TokenResponse {
	return models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    time.Until(expiresAt).Milliseconds() / 1000,
		ExpiresAt:    expiresAt,
	}
}
