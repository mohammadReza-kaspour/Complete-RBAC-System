package inbound

type RoleView struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PermissionView struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Resource    string `json:"resource"`
	Action      string `json:"action"`
}

type ListQuery struct {
	Search   string `json:"search"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}

func (q *ListQuery) Normalized() {
	n := q
	if n.Page < 1 {
		n.Page = 1
	}
	if n.PageSize < 1 {
		n.PageSize = 20
	}
	if n.PageSize > 100 {
		n.PageSize = 100
	}
}

type Page[T any] struct {
	Items      []T `json:"items"`
	Total      int `json:"total"`
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	TotalPages int `json:"totalPages"`
}
