package otp

import (
	"crypto/subtle"
	"time"
	"uuid"

	"github.com/mohammadReza-kaspour/Complete-RBAC-System/internal/domain/domainerr"
)

const (
	OtpMaxAttempts  = 5
	OtpPurposeLogin = "login"
)

type OTP struct {
	Id        string
	UserId    string
	Code      string
	Purpose   string
	ExpiresAt time.Time
	Consumed  bool
	Attempts  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(userId string, code string, purpose string, ttl time.Duration) (*OTP, error) {
	now := time.Now()
	otp := &OTP{
		Id:        uuid.NewV7().String(),
		UserId:    userId,
		Code:      code,
		Purpose:   purpose,
		ExpiresAt: now.Add(ttl),
		Consumed:  false,
		Attempts:  0,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := otp.Validate(); err != nil {
		return nil, err
	}
	return otp, nil
}
func (o *OTP) Validate() error {
	return nil
}
func (o *OTP) IsExpired(now time.Time) bool {
	return now.After(o.ExpiresAt)
}
func (o *OTP) IsMatched(code string) bool {
	return subtle.ConstantTimeCompare([]byte(o.Code), []byte(code)) == 1
}
func (o *OTP) RegisterFaildAttempt() {
	o.Attempts++
}
func (o *OTP) Consume() {
	o.Consumed = true
}
func (o *OTP) Verify(code string) error {
	now := time.Now()

	if o.Consumed {
		return domainerr.New(domainerr.OTPAlreadyUsed, "otp already used")
	}

	if o.IsExpired(now) {
		return domainerr.New(domainerr.OTPExpired, "otp expired")
	}

	if o.Attempts > OtpMaxAttempts {
		return domainerr.New(domainerr.OTPTooManyAttempts, "too many attempts")
	}

	if o.IsMatched(code) {
		return domainerr.New(domainerr.OTPInvalidCode, "invalid code")
	}

	return nil
}
