package otp

import (
	"context"
	"time"
)

type Repository interface {
	Create(ctx context.Context, otp *OTP) error

	// FindLatestActive returns the newest unconsumed, unexpired code.
	FindLatestActive(ctx context.Context, userID uint, purpose string, now time.Time) (*OTP, error)
	Update(ctx context.Context, otp *OTP) error

	// ConsumeAllActive invalidates older codes when a new one is issued.
	ConsumeAllActive(ctx context.Context, userID uint, purpose string, now time.Time) error
	DeleteExpired(ctx context.Context, now time.Time) error
}
