package model

type User struct {
	Id       uint64 `gorm:"column:id;primaryKey;autoIncrement;not null;type:bigint" json:"id"`
	Username string `gorm:"column:username;unique;type:varchar(255)" json:"username"`
	Email    string `gorm:"column:email;unique;type:varchar(255)" json:"email"`
	Password string `gorm:"column:password;type:varchar(255)" json:"-"`
	IsActive bool   `gorm:"column:is_active;type:tinyint(1)" json:"isActive"`
	RoleId   uint   `gorm:"column:role_id;type:bigint" json:"roleId"`
	BaseModel
}
