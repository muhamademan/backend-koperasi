package request

// Response struct TRUE ketika CREATE USER
type UserCreateRequest struct {
	NIK      string `json:"nik" form:"nik" validate:"required" gorm:"not null"`
	Name     string `json:"name" form:"name" validate:"required" gorm:"not null"`
	Email    string `json:"email" form:"email" validate:"required,email" gorm:"not null"` // must valid email
	Address  string `json:"address" form:"address" gorm:"not null"`
	Phone    string `json:"phone" form:"phone" validate:"lte=12" gorm:"not null"`
	Password string `json:"password" form:"password" validate:"required,min=6" gorm:"not null"` // min 6 char
	Role     string `json:"role" form:"role" gorm:"not null"`
}

// Response struct TRUE ketika UPDATE USER
type UserUpdateRequest struct {
	NIK     string `json:"nik" form:"nik" validate:"required" gorm:"not null"`
	Name    string `json:"name" form:"name" validate:"required" gorm:"not null"`
	Email   string `json:"email" form:"email" validate:"required,email" gorm:"not null"` // must valid email
	Address string `json:"address" form:"address"`
	Phone   string `json:"phone" form:"phone" validate:"lte=13"` // max 13 number phone
	Role    string `json:"role" form:"role" gorm:"not null"`
}
