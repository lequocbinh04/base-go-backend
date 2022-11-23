package profilebiz

type profileStorage interface {
}

type profileBiz struct {
	store profileStorage
}

func NewProfileBiz(store profileStorage) *profileBiz {
	return &profileBiz{
		store: store,
	}
}

func (biz *profileBiz) CreateProfile() error {
	return nil
}
