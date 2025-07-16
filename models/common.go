package models

// Claims — JWT token ma'lumotlari
type Claims struct {
	User_id   string `json:"user_id"`
	User_role string `json:"user_role"`
}

// OtpData — OTP yuborish uchun
type OtpData struct {
	Otp   string `json:"otp"`
	Email string `json:"email"`
}

// CheckOtpResp — OTP tekshiruvi javobi
type CheckOtpResp struct {
	Is_right string `json:"is_right"`
}

// AuthResp — login/verify javobi
type AuthResp struct {
	Access_token string `json:"access_token"`
}

// CheckExists — mavjudlikni tekshirish
type CheckExists struct {
	Status    string `json:"status"`
	Is_exists bool   `json:"is_exists"`
}

// GetById — ID orqali ma’lumot olish
type GetById struct {
	Id string `json:"id"`
}

// Common — umumiy mavjudlik tekshiruvi
type Common struct {
	Table_name  string `json:"table_name"`
	Column_name string `json:"column_name"`
	Expvalue    string `json:"expvalue"`
}

// CommonResp — umumiy mavjudlik javobi
type CommonResp struct {
	IsExists bool `json:"is_exists"`
}



type Check_User struct {
	Email string `json:"email"`
}


type LoginReq struct {
	Email         string `json:"email"`
	User_password string `json:"password"`
}