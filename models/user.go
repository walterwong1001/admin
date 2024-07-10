package models

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" binding:"required"` // 用户名
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"phone" binding:"required"`
}

type Profile struct {
	ID       int64 // 关联User ID
	Nickname string
	Avatar   string
	Motto    string // 个性签名，座右铭，箴言
}
