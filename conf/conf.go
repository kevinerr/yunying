package conf

import (
	"gopkg.in/ini.v1"
	"strings"
	"yy/model"
	"yy/pkg/util"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

// Init 初始化配置
func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		util.Logger().Info("配置文件读取错误，请检查文件路径:", err)
		panic(err)
	}
	LoadMysqlData(file)
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true&loc=Local"}, "")
	model.Database(path)
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String() + "#123456abAB"
	DbName = file.Section("mysql").Key("DbName").String()
}
