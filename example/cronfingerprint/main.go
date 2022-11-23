package main

import (
	"cronbrowser/appCommon"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	goservice "github.com/lequocbinh04/go-sdk"
	"github.com/lequocbinh04/go-sdk/plugin/storage/sdkgorm"
	"gorm.io/gorm"
)

type respData struct {
	OsType         string `json:"os_type"`
	OsVersion      string `json:"os_version"`
	Webgl          string `json:"webgl"`
	Screen         string `json:"screen"`
	BrowserType    string `json:"browser_type"`
	BrowserVersion string `json:"browser_version"`
	UserAgent      string `json:"useragent"`
}

func main() {
	service := goservice.New(
		goservice.WithName("cronbrowser"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main", appCommon.DBMain)),
	)

	service.Init()

	totalPage := 20216
	for i := 0; i < totalPage; i++ {
		client := resty.New()
		resp, err := client.R().
			EnableTrace().
			Get("https://app.undetectable.io/configs/store-json?draw=1&columns%5B0%5D%5Bdata%5D=&columns%5B0%5D%5Bname%5D=&columns%5B0%5D%5Bsearchable%5D=true&columns%5B0%5D%5Borderable%5D=false&columns%5B0%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B0%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B1%5D%5Bdata%5D=id&columns%5B1%5D%5Bname%5D=&columns%5B1%5D%5Bsearchable%5D=true&columns%5B1%5D%5Borderable%5D=true&columns%5B1%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B1%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B2%5D%5Bdata%5D=os_type&columns%5B2%5D%5Bname%5D=&columns%5B2%5D%5Bsearchable%5D=true&columns%5B2%5D%5Borderable%5D=true&columns%5B2%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B2%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B3%5D%5Bdata%5D=browser_type&columns%5B3%5D%5Bname%5D=&columns%5B3%5D%5Bsearchable%5D=true&columns%5B3%5D%5Borderable%5D=true&columns%5B3%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B3%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B4%5D%5Bdata%5D=useragent&columns%5B4%5D%5Bname%5D=&columns%5B4%5D%5Bsearchable%5D=true&columns%5B4%5D%5Borderable%5D=false&columns%5B4%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B4%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B5%5D%5Bdata%5D=webgl&columns%5B5%5D%5Bname%5D=&columns%5B5%5D%5Bsearchable%5D=true&columns%5B5%5D%5Borderable%5D=false&columns%5B5%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B5%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B6%5D%5Bdata%5D=screen&columns%5B6%5D%5Bname%5D=&columns%5B6%5D%5Bsearchable%5D=true&columns%5B6%5D%5Borderable%5D=false&columns%5B6%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B6%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B7%5D%5Bdata%5D=hardware_concurrency&columns%5B7%5D%5Bname%5D=&columns%5B7%5D%5Bsearchable%5D=true&columns%5B7%5D%5Borderable%5D=true&columns%5B7%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B7%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B8%5D%5Bdata%5D=device_memory&columns%5B8%5D%5Bname%5D=&columns%5B8%5D%5Bsearchable%5D=true&columns%5B8%5D%5Borderable%5D=true&columns%5B8%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B8%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B9%5D%5Bdata%5D=created_at&columns%5B9%5D%5Bname%5D=&columns%5B9%5D%5Bsearchable%5D=true&columns%5B9%5D%5Borderable%5D=true&columns%5B9%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B9%5D%5Bsearch%5D%5Bregex%5D=false&columns%5B10%5D%5Bdata%5D=browser_version&columns%5B10%5D%5Bname%5D=&columns%5B10%5D%5Bsearchable%5D=true&columns%5B10%5D%5Borderable%5D=false&columns%5B10%5D%5Bsearch%5D%5Bvalue%5D=&columns%5B10%5D%5Bsearch%5D%5Bregex%5D=false&order%5B0%5D%5Bcolumn%5D=1&order%5B0%5D%5Bdir%5D=desc&start=" + fmt.Sprintf("%d", i*100) + "&length=100&search%5Bvalue%5D=&search%5Bregex%5D=false&_=1667131412261")
		if err != nil {
			continue
		}
		data := make(map[string]interface{})
		body := string(resp.Body())
		if json.Unmarshal([]byte(body), &data) != nil {
			continue
		}
		res, ok := data["data"].([]interface{})
		if !ok {
			continue
		}
		resStr, err := json.Marshal(res)
		if err != nil {
			continue
		}
		var respData []respData
		if json.Unmarshal(resStr, &respData) != nil {
			continue
		}
		for _, v := range respData {
			db := service.MustGet(appCommon.DBMain).(*gorm.DB)
			if db.Table("fingerprint").Create(map[string]interface{}{
				"os_type":         v.OsType,
				"os_version":      v.OsVersion,
				"browser_type":    v.BrowserType,
				"browser_version": v.BrowserVersion,
				"user_agent":      v.UserAgent,
				"webgl":           v.Webgl,
				"screen":          v.Screen,
			}).Error != nil {
				continue
			}
		}
	}
}
