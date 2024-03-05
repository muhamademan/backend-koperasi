package request

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email" gorm:"not null"`
	Password string `json:"password" form:"password" validate:"required" gorm:"not null"`
}
