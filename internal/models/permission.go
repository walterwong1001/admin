package models

type Allowed uint8

const (
	YES Allowed = 1
	NO  Allowed = 0
)

type Permission struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Path        string  `json:"path" binding:"required"`
	Method      string  `json:"method" binding:"required,http_method"`
	Allowed     Allowed `json:"allowed"`
	Description string  `json:"description"`
	CreateTime  int64   `json:"create_time"`
}

func (p *Permission) IsAllowed() bool {
	return p.Allowed == YES
}

type PermissionFilter struct{}
