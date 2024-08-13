package models

type Permission struct {
	ID          uint64
	Name        string `json:"name" binding:"required"`
	Path        string `json:"path" binding:"required"`
	Method      string `json:"method" binding:"required,http_method"`
	Description string
	CreateTime  int64
}
