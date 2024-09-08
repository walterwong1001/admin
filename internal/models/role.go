package models

type Role struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
	CreateTime  int64  `json:"create_time"`
}

type RoleFilter struct {
	Name  string `form:"name"`
	Start int64  `form:"start"`
	End   int64  `form:"end"`
}
