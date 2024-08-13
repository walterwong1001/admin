package models

type Navigation struct {
	ID     uint64
	PID    uint64 `json:"pid" gorm:"column:pid"`
	Title  string `json:"title" binding:"required"`
	URI    string
	Icon   string
	Header string
	Seq    uint8
}
