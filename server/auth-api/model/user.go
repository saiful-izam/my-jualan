package model

type User struct {
	Id       uint64 `gorm:"column:id, primaryKey, autoIncrement, not null" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Email    string `gorm:"column:email, unique" json:"email"`
	Password string `gorm:"column:password" json:"-"`
	IsActive bool   `gorm:"column:is_active" json:"isActive"`
	RoleId   uint   `gorm:"column:role_id" json:"roleId"`
	BaseModel
}
