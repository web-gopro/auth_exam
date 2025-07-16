package models

import "time"

// User — oddiy foydalanuvchi (client)
type User struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"` // "active", "deleted"
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"` // Parol JSON response’da ko‘rinmaydi
	CreatedAt time.Time `json:"created_at"`
	CreatedBy *string   `json:"created_by,omitempty"` // Sysuser ID (admin tomonidan yaratilgan bo‘lsa)
}

type UserCreReq struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"` // "active", "deleted"
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"` // Parol JSON response’da ko‘rinmaydi
	CreatedAt time.Time `json:"created_at"`
	CreatedBy *string   `json:"created_by,omitempty"` // Sysuser ID (admin tomonidan yaratilgan bo‘lsa)
	Otp       string    `json:"otp"`
}

// UserLogin — foydalanuvchi login uchun



// UserCreateResp — foydalanuvchi yaratish javobi
type UserCreateResp struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}