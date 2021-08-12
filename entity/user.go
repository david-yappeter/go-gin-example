package entity

type User struct {
	UuID     string `json:"uu_id" gorm:"type:varchar(255);primaryKey"`
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);not null"`
	Password string `json:"-" gorm:"type:varchar(255);not null"`
}

type NewUserForm struct {
	UuID     string `json:"uu_id"`
	Name     string `json:"name" binding:"required,min=1,max=255"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=4,max=255"`
}

type UpdateUserForm struct {
	UuID string `json:"uu_id" binding:"required,min=36,max=36"`
	Name string `json:"name" binding:"required,min=1,max=255"`
}

type UserLoginForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" bindingn:"required,min=4,max=255"`
}

func (u *User) TableName() string {
	return "user"
}
