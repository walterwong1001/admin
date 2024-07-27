package models

type Navigation struct {
	ID     uint64
	PID    uint64 `json:"pid" gorm:"column:pid"`
	Title  string
	URI    string
	Icon   string
	Header string
	Seq    uint8
}
