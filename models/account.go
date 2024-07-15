package models

type AccountType string

const (
	AccountTypeUsername AccountType = "username"
	AccountTypeEmail    AccountType = "email"
	AccountTypeMobile   AccountType = "phone"
	AccountActivity     uint        = 1
	AccountLocked       uint        = 0
)

type Account struct {
	ID         uint64      `json:"id"`
	UserID     uint64      `json:"user_id"`    // 外键，关联 User
	Identifier string      `json:"identifier"` // 可以是手机号、邮箱或其他用户名
	Password   string      `json:"password"`
	Type       AccountType `json:"type"`
	Status     uint        `json:"status"` // 账户状态（例如激活、锁定等）
	CreateTime int64       `json:"create_time"`
}
