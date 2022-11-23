package appCommon

import "time"

type SQLModel struct {
	Id        int64      `json:"-" gorm:"column:id"`
	FakeId    *UID       `json:"id" gorm:"-"`
	Status    string     `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (sqlModel *SQLModel) Mask(dbType DbType) {
	uid := NewUID(uint32(sqlModel.Id), int(dbType), 1)
	sqlModel.FakeId = &uid
}

func (sqlModel *SQLModel) PrepareForInsert() {
	now := time.Now().UTC()
	sqlModel.Id = 0
	sqlModel.CreatedAt = &now
	sqlModel.UpdatedAt = &now
}

func (sqlModel *SQLModel) PrepareForUpdate() {
	now := time.Now().UTC()
	sqlModel.UpdatedAt = &now
}
