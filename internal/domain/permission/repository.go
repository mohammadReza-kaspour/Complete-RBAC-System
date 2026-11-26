package permission

import "context"

type ListFilter struct {
	Search   string
	Page     int
	PageSize int
}

type Repository interface {
	Create(ctx context.Context, permission *Permission) error
	FindByID(ctx context.Context, id uint) (*Permission, error)
	FindByName(ctx context.Context, name string) (*Permission, error)
	List(ctx context.Context, listFilter ListFilter) ([]Permission, int64, error)
	Update(ctx context.Context, permission *Permission) error
	Delete(ctx context.Context, id uint) error
	Count(ctx context.Context) (int64, error)
}
