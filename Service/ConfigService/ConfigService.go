package ConfigService

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type DbInfo struct {
	TYPE      string
	USER      string
	PASSWORD  string
	HOST      string
	NAME      string
	HTTP_PORT string
}

func GetAppConfig(key string) *DbInfo {
	cfg, err := ini.Load("Config/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	d := new(DbInfo)
	_ = cfg.Section(key).MapTo(d)
	return d
}

func GetServerConfig() *DbInfo {
	return GetAppConfig("server")
}

func GetDbConfig() *DbInfo {
	return GetAppConfig("database")
}
