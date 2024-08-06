package models

type Permission struct {
	ID          uint64
	Name        string
	Path        string
	Method      string `json:"method" binding:"required,method"`
	Description string
	CreateTime  int64
}
