package fingerprintstorage

import (
	"context"
	"cronbrowser/appCommon"
	fingerprintmodel "cronbrowser/module/fingerprint/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindFingerprint(ctx context.Context, condition map[string]interface{}, filter *fingerprintmodel.Filter) (*fingerprintmodel.Fingerprint, error) {
	db := s.db

	var data fingerprintmodel.Fingerprint

	if v := filter.OsType; v != "" {
		db = db.Where("os_type = ?", v)
	}
	if v := filter.OsVersion; v != "" {
		db = db.Where("os_version = ?", v)
	}
	if v := filter.BrowserType; v != "" {
		db = db.Where("browser_type = ?", v)
	}
	if v := filter.BrowserVersion; v != "" {
		db = db.Where("browser_version = ?", v)
	}
	if v := filter.Webgl; v != "" {
		db = db.Where("webgl = ?", v)
	}
	if v := filter.Screen; v != "" {
		db = db.Where("screen = ?", v)
	}

	db = db.Where(condition).Order("random()")

	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, appCommon.ErrRecordNotFound
		}
		return nil, appCommon.ErrDB(err)
	}
	return &data, nil
}
