package models

type Allowed uint8

const (
	YES Allowed = 1
	NO  Allowed = 0
)

type Permission struct {
	ID          uint64
	Name        string `json:"name" binding:"required"`
	Path        string `json:"path" binding:"required"`
	Method      string `json:"method" binding:"required,http_method"`
	Allowed     Allowed
	Description string
	CreateTime  int64
}

func (p *Permission) IsAllowed() bool {
	return p.Allowed == YES
}
