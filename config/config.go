package config

import (
	"encoding/json"
	"os"
)

/**
mysql
*/
type DataBase struct {
	Drive string `json:"drive"`
	Post  string `json:"post"`
	User  string `json:"user"`
	Pwd   string `json:"pwd"`
	Database string `json:"database"`

} 

type AppConfig struct {
	AppName    string `json:"app_name"`
	Port       string `json:"port"`
	StaticPath string `json:"static_path"`
	Mode       string `json:"mode"`
	DataBase   DataBase `json:"data_base"`
	Redis Redis	`json:"redis"`
}


/**
 * Redis 配置
 */
type Redis struct {
	NetWork  string `json:"net_work"`
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Prefix   string `json:"prefix"`
}


//var ServConfig AppConfig

func InitConfig() *AppConfig  {

	var config *AppConfig
	file, err := os.Open("/Users/ge/Documents/code/gopath/goadmin/config/config.json")
	if err!=nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	//config := AppConfig{}
	err = decoder.Decode(&config)
	if err!=nil {
		panic(err.Error())
	}
	return config
}