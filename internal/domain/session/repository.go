package session

import (
	"context"
	"time"
)

type Repository interface {
	Create(ctx context.Context, session *Session) error
	FindByRefreshToken(ctx context.Context, token string) (*Session, error)
	Revoke(ctx context.Context, id uint, now time.Time) error
	RevokeAllForUser(ctx context.Context, userID uint, now time.Time) error
	CountActive(ctx context.Context, now time.Time) (int64, error)
	DeleteExpired(ctx context.Context, now time.Time) error
}
