package outbound

import (
	"context"
	"time"
)

type Claims struct {
	UserId    string
	Mobile    string
	Roles     []string
	ExpiresAt time.Time
}

type TokenService interface {
	GenerateAccessToken(ctx context.Context, claim *Claims) (token string, expiresIn time.Duration, err error)
	ValidateAccessToken(ctx context.Context, token string) (*Claims, error)

	GenerateRefreshToken(ctx context.Context, claim *Claims) (token string, expiresIn time.Duration, err error)
	ValidateRefreshToken(ctx context.Context, token string) (*Claims, error)
}
