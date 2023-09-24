package utils

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	AppMode      string
	HttpPort     string
	Db           string
	JwtKey       string
	DbHost       string
	DbPort       string
	DbName       string
	DbUser       string
	DbPassWord   string
	RobotWebhook string
	RobotCode    string
	AgentId      string
	MiniAppId    string
	AppKey       string
	AppSecret    string
)

// GetAppPath 获取编译后后的路径地址
func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}
func init() {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	path = path[:index]

	conf, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatal("配置文件读取错误，请检查文件路径 ", err)
	}
	loadServer(conf)
	LoadData(conf)
	LoadRobot(conf)
}

// 初始化项目配置信息
func loadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString(":j00k07w19")
}

// LoadData 初始化数据库相关配置
func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")

}

// LoadRobot 初始化钉钉机器人配置信息
func LoadRobot(file *ini.File) {
	RobotWebhook = file.Section("robot").Key("Webhook").MustString("https://oapi.dingtalk.com/robot/send?access_token=7c87c8d6899ca26e24b08ec3a42eb7bdc5bc405e9868a6fa3f2120c1a91ac298")
	RobotCode = file.Section("robot").Key("RobotCode").MustString("dingybxvtoxyjto2rrpj")
	AgentId = file.Section("robot").Key("AgentId").MustString("2741681097")
	MiniAppId = file.Section("robot").Key("MiniAppId").MustString("5000000004901971")
	AppKey = file.Section("robot").Key("AppKey").MustString("dingybxvtoxyjto2rrpj")
	AppSecret = file.Section("robot").Key("AppSecret").MustString("tMsASWsJNXmbVWPVpIgQ89H0rGSR1l_DVtZa3qQS8g_KBTh3Tte_V6VdK_lZfMks")

}
