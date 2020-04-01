package setting

import (
	"GoWebTrader/util"
	"github.com/mitchellh/go-homedir"
)

type WebTraderSetting struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}

var WebTrader *WebTraderSetting

func init() {
	WebTrader = &WebTraderSetting{}

	dir, _ := homedir.Dir()
	filename := dir + "/.vntrader/web_trader.json"

	JsonParse := util.NewJsonStruct()
	JsonParse.Load(filename, WebTrader)

}
