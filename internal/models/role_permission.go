package models

type RolePermission struct {
	RoleId       uint64 `json:"role_id"`
	PermissionId uint64 `json:"permission_id"`
}
