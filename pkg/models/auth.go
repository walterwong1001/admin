package models

// AuthCredential 认证凭证
type AuthCredential struct {
	Type       AccountType `json:"type" binding:"required"`       // 登录类型
	Identifier string      `json:"identifier" binding:"required"` // 用户名、邮箱或手机号
	Secret     string      `json:"secret" binding:"required"`     // 密码或短信验证码
}

// AuthPrinciple 认证后主体，用于返回
type AuthPrinciple struct {
	Token string `json:"token"`
}
