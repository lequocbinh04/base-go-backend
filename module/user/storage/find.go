package userstorage

import (
	"context"
	"cronbrowser/appCommon"
	"cronbrowser/module/user/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, appCommon.ErrRecordNotFound
		}

		return nil, appCommon.ErrDB(err)
	}

	return &user, nil
}
