package model

type Role struct {
	Id       uint64 `gorm:"column:id;primaryKey;autoIncrement;not null;type:bigint" json:"id"`
	RoleName string `gorm:"column:role_name;type:varchar(50)" json:"roleName"`
	IsActive bool   `gorm:"column:is_active;type:bit(1)" json:"isActive"`
	BaseModel
}
