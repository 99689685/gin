package utils

import "fmt"

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
	file, err := ini.Load("../config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请坚持文件路径")
	}
	loadServer(file)
	LoadData(file)
}
func loadServer(file *ini.File) {

}
func LoadData(file *ini.File) {

}
