package profilemodel

import (
	"cronbrowser/appCommon"
	"cronbrowser/module/user/model"
	"time"
)

const EntityName = "Profile"

type Profile struct {
	appCommon.SQLModel
	Name        string                 `json:"name" gorm:"column:name"`
	Information map[string]interface{} `json:"information" gorm:"column:information"`
	ExpiredAt   *time.Time             `json:"expired_at" gorm:"column:expired_at"`
	User        usermodel.User         `json:"user"`
}

func (Profile) Table() string {
	return "profiles"
}

type ProfileCreate struct {
	Name        string                 `json:"name" gorm:"column:name"`
	Information map[string]interface{} `json:"information" gorm:"column:information"`
}
