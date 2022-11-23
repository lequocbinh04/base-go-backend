package fingerprintmodel

type Filter struct {
	OsType         string `json:"os_type" gorm:"column:os_type"`
	OsVersion      string `json:"os_version" gorm:"column:os_version"`
	BrowserType    string `json:"browser_type" gorm:"column:browser_type"`
	BrowserVersion string `json:"browser_version" gorm:"column:browser_version"`
	Webgl          string `json:"webgl" gorm:"column:webgl"`
	Screen         string `json:"screen" gorm:"column:screen"`
}
