package models

type AccountType uint8

const (
	AccountTypeUsername AccountType = iota + 1
	AccountTypeEmail
	AccountTypeMobile
)

const (
	AccountActivity uint8 = 1
	AccountLocked   uint8 = 0
)

type Account struct {
	ID         uint64      `json:"id"`
	UserID     uint64      `json:"user_id"`                       // 外键，关联 User
	Identifier string      `json:"identifier" binding:"required"` // 可以是手机号、邮箱或其他用户名
	Password   string      `json:"password" binding:"required"`
	Type       AccountType `json:"type" binding:"required"`
	Status     uint8       `json:"status"` // 账户状态（例如激活、锁定等）
	CreateTime int64       `json:"create_time"`
}
