package permission

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"uuid"

	"github.com/mohammadReza-kaspour/Complete-RBAC-System/internal/domain/domainerr"
)

var (
	resourceRe = regexp.MustCompile(`^[a-z][a-z0-9_]{1,63}$`)
	actionRe   = regexp.MustCompile(`^[a-z][a-z0-9_]{1,63}$`)
)

type Permission struct {
	Id          string
	Name        string
	Description string
	Resource    string
	Action      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func New(resource string, action string, description string) (*Permission, error) {
	now := time.Now()
	permission := &Permission{
		Id:          uuid.NewV7().String(),
		Description: strings.TrimSpace(description),
		Resource:    strings.ToLower(strings.TrimSpace(resource)),
		Action:      strings.ToLower(strings.TrimSpace(action)),
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := permission.Validate(); err != nil {
		return nil, err
	}

	permission.setName()
	return permission, nil
}

func (p *Permission) setName() {
	p.Name = fmt.Sprintf("%s:%s", p.Resource, p.Action)
}
func (p *Permission) Validate() error {
	if !resourceRe.MatchString(p.Resource) || !actionRe.MatchString(p.Action) {
		return domainerr.New(domainerr.InvalidPermissionParts, "invalid permission parts")
	}
	return nil
}
func (p *Permission) SetDescription(description string) {
	p.Description = strings.TrimSpace(description)
}
func (p *Permission) SetParts(resource string, action string) error {
	p.Resource = strings.ToLower(strings.TrimSpace(resource))
	p.Action = strings.ToLower(strings.TrimSpace(action))
	if err := p.Validate(); err != nil {
		return err
	}

	p.setName()
	return nil
}
