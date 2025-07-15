package models

import "time"

// SysUser — admin panel foydalanuvchisi
type SysUser struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"` // "active", "deleted"
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy *string   `json:"created_by,omitempty"`
}

type SysUserCretReq struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"` // "active", "deleted"
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy *string   `json:"created_by,omitempty"`
	Role      string    `json:"role"`
}

// Role — tizimdagi rollar
type Role struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"` // "active", "deleted"
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy *string   `json:"created_by,omitempty"`
}

// SysUserRole — sysuser va role orasidagi bog‘lovchi jadval
type SysUserRole struct {
	ID        string `json:"id"`
	SysUserID string `json:"sysuser_id"`
	RoleID    string `json:"role_id"`
}

// UserCreateResp — foydalanuvchi yaratish javobi
type SysUserCreateResp struct {
	Id     string `json:"id"`
	Role string `json:"role"`
}