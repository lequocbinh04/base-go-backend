package fingerprintbiz

import (
	"context"
	"cronbrowser/appCommon"
	fingerprintmodel "cronbrowser/module/fingerprint/model"
)

type FindRandomFingerprintStore interface {
	FindFingerprint(ctx context.Context, condition map[string]interface{}, filter *fingerprintmodel.Filter) (*fingerprintmodel.Fingerprint, error)
}

type findRandomFingerprintBiz struct {
	store FindRandomFingerprintStore
}

func NewFindRandomFingerprintBiz(store FindRandomFingerprintStore) *findRandomFingerprintBiz {
	return &findRandomFingerprintBiz{store: store}
}

func (biz *findRandomFingerprintBiz) FindRandomFingerprint(ctx context.Context, filter *fingerprintmodel.Filter) (*fingerprintmodel.Fingerprint, error) {
	data, err := biz.store.FindFingerprint(ctx, nil, filter)
	if err != nil {
		return nil, appCommon.ErrCannotGetEntity(fingerprintmodel.EntityName, err)
	}
	return data, nil
}
