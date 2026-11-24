package user

import (
	"strings"
	"time"
	"uuid"

	"github.com/mohammadReza-kaspour/Complete-RBAC-System/internal/domain/role"
)

type User struct {
	Id        string
	UserName  string
	Email     string
	Mobile    string
	IsActive  bool
	Roles     []role.Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(userName string, email string, mobile string) (*User, error) {
	now := time.Now()
	user := &User{
		Id:        uuid.NewV7().String(),
		UserName:  userName,
		Email:     email,
		Mobile:    mobile,
		IsActive:  true,
		CreatedAt: now,
		UpdatedAt: now,
		Roles:     []role.Role{},
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Validate() error {
	if err := validateUserName(u.UserName); err != nil {
		return err
	}
	if err := validateEmail(u.Email); err != nil {
		return err
	}
	if err := validateMobile(u.Mobile); err != nil {
		return err
	}
	return nil
}
func (u *User) Rename(userName string) error {
	userName = strings.TrimSpace(userName)
	if err := validateUserName(userName); err != nil {
		return err
	}

	u.UserName = userName
	return nil
}
func (u *User) ChangeEmail(email string) error {
	email = strings.ToLower(strings.TrimSpace(email))
	if err := validateEmail(email); err != nil {
		return err
	}

	u.Email = email
	return nil
}
func (u *User) ChangeMobile(mobile string) error {
	mobile = strings.TrimSpace(mobile)
	if err := validateMobile(mobile); err != nil {
		return err
	}

	u.Mobile = mobile
	return nil
}
func (u *User) Activate() {
	u.IsActive = true
}
func (u *User) Deactivate() {
	u.IsActive = false
}
func (u *User) HasRole(name string) bool {
	for _, role := range u.Roles {
		if role.Name == name {
			return true
		}
	}
	return false
}
