package usermodel

import "cronbrowser/appCommon"

const (
	StatusActive = "active"
	StatusBan    = "ban"
)

type User struct {
	appCommon.SQLModel
	Name           string  `json:"name" gorm:"column:name;"`
	Email          string  `json:"email" gorm:"column:email"`
	Password       string  `json:"-" gorm:"column:password"`
	CrondataToken  string  `json:"-" gorm:"column:crondata_token"`
	Balance        float64 `json:"balance" gorm:"column:balance"`
	Role           string  `json:"role" gorm:"column:role"`
	BalanceProfile float64 `json:"balance_profile" gorm:"column:balance_profile"`
}

func (User) TableName() string {
	return "users"
}
