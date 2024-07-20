package models

type AuthCredential struct {
	Type       AccountType `json:"type"`       // 登录类型
	Identifier string      `json:"identifier"` // 用户名、邮箱或手机号
	Secret     string      `json:"secret"`     // 密码或短信验证码
}

type AuthPrinciple struct {
	Token string `json:"token"`
}
