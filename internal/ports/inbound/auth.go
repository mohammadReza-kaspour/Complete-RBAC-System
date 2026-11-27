package inbound

import (
	"context"
	"time"
)

type LoginCommand struct {
	Mobile string `json:"mobile"`
}
type VerifyOtpCommand struct {
	Mobile    string `json:"mobile"`
	Code      string `json:"code"`
	Ip        string `json:"-"`
	UserAgent string `json:"-"`
}
type RefreshCommand struct {
	RefreshToken string `json:"refreshToken"`
	Ip           string `json:"-"`
	UserAgent    string `json:"-"`
}
type LogoutCommand struct {
	RefreshToken string `json:"refreshToken"`
}

type AuthTokens struct {
	AccessToken  string `json:"accessToken"`
	ExpiresIn    int64  `json:"expiresIn"`
	RefreshToken string `json:"refreshToken"`
}

type AuthenticatedUser struct {
	Id        string    `json:"id"`
	UserName  string    `json:"userName"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	IsActive  bool      `json:"isActive"`
	Roles     []string  `json:"roles"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VerifyResult struct {
	Tokens AuthTokens        `json:"tokens"`
	User   AuthenticatedUser `json:"user"`
}

type AuthManagement interface {
	Login(ctx context.Context, cmd LoginCommand) error
	VerifyOtp(ctx context.Context, cmd VerifyOtpCommand) (*VerifyResult, error)
	Refresh(ctx context.Context, cmd RefreshCommand) (*AuthTokens, error)
	Logout(ctx context.Context, cmd LogoutCommand) error
}
