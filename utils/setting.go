package utils

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassWord string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatal("配置文件读取错误，请检查文件路径 ", err)
	}
	loadServer(file)
	LoadData(file)
}

func loadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}
func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")

}
