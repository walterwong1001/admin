package models

type Role struct {
	ID          uint64
	Name        string `json:"name" binding:"required"`
	Description string
	CreateTime  uint64
}
