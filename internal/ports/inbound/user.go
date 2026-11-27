package inbound

import (
	"context"
	"time"
)

type CreateUserInput struct {
	UserName string   `json:"userName"`
	Email    string   `json:"email"`
	Mobile   string   `json:"mobile"`
	RoleIds  []string `json:"roleIds"`
}
type UpdateUserInput struct {
	UserName *string `json:"userName"`
	Email    *string `json:"email"`
	Mobile   *string `json:"mobile"`
	IsActive *bool   `json:"isActive"`
}
type UserView struct {
	Id        string     `json:"id"`
	UserName  string     `json:"userName"`
	Email     string     `json:"email"`
	Mobile    string     `json:"mobile"`
	IsActive  bool       `json:"isActive"`
	Roles     []RoleView `json:"roles"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type UserManager interface {
	Create(ctx context.Context, cmd CreateUserInput) (*UserView, error)
	Get(ctx context.Context, id string) (*UserView, error)
	List(ctx context.Context, listQuery *ListQuery) (*Page[UserView], error)
	Update(ctx context.Context, id string, cmd UpdateUserInput) (*UserView, error)
	Delete(ctx context.Context, id string) error
}
