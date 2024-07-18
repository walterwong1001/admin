package models

type User struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name" binding:"required"` // 用户名
	Email      string `json:"email,omitempty" binding:"required,email"`
	Mobile     string `json:"mobile,omitempty" binding:"required"`
	Password   string `json:"password,omitempty" binding:"required" gorm:"-"`
	CreateTime int64  `json:"create_time"`
}

type Profile struct {
	ID       uint64 `json:"machine"` // 关联User ID
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Motto    string `json:"motto"` // 个性签名，座右铭，箴言
}
