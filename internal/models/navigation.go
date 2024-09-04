package models

type Navigation struct {
	ID        uint64 `json:"id"`
	PID       uint64 `json:"pid,omitempty" gorm:"column:pid"`
	Name      string `json:"name" binding:"required"`
	Path      string `json:"path,omitempty" binding:"required"`
	Icon      string `json:"icon,omitempty"`
	Component string `json:"component,omitempty"`
	Redirect  string `json:"redirect,omitempty"`
	Header    string `json:"header,omitempty"`
	Seq       uint8  `json:"seq"`
}
