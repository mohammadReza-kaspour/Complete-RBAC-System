package session

import (
	"time"
	"uuid"
)

type Session struct {
	Id               string
	UserID           string
	RefreshTokenHash string
	UserAgent        string // device info captured at login, "Mozilla/5.0 (Windows NT 10.0)"
	IP               string
	ExpiresAt        time.Time
	RevokedAt        *time.Time
	CreatedAt        time.Time
}

func New(userID string, refreshToken, ip, userAgent string, ttl time.Duration) (*Session, error) {
	now := time.Now()
	session := &Session{
		Id:               uuid.NewV7().String(),
		UserID:           userID,
		RefreshTokenHash: refreshToken,
		UserAgent:        userAgent,
		IP:               ip,
		ExpiresAt:        now.Add(ttl),
		CreatedAt:        now,
	}
	if err := session.Validate(); err != nil {
		return nil, err
	}

	return session, nil
}
func (s *Session) Validate() error {
	return nil
}

func (s *Session) IsActive(now time.Time) bool {
	return s.RevokedAt == nil && now.Before(s.ExpiresAt)
}
func (s *Session) Revoke(now time.Time) {
	if s.RevokedAt == nil {
		t := now
		s.RevokedAt = &t
	}
}
