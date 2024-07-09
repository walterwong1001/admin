package models

type AccountType string

const (
	AccountTypeEmail    AccountType = "email"
	AccountTypePhone    AccountType = "phone"
	AccountTypeUsername AccountType = "username"
)

type Account struct {
	ID         uint        `json:"id"`
	UserID     uint        `json:"user_id"`    // 外键，关联 User
	Identifier string      `json:"identifier"` // 可以是手机号、邮箱或其他用户名
	Password   string      `json:"password"`
	Type       AccountType `json:"type"`
	Status     string      `json:"status"` // 账户状态（例如激活、锁定等）
}
