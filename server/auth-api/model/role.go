package model

type Role struct {
	Id       uint64 `gorm:"column:id, primaryKey, autoIncrement, not null" json:"id"`
	RoleName string `gorm:"column:role_name" json:"roleName"`
	IsActive bool   `gorm:"column:is_active" json:"isActive"`
	BaseModel
}
