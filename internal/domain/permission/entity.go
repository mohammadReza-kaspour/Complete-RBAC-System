package permission

import (
	"time"
)

type Permission struct {
	Id          string
	Name        string
	Description string
	Code        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
