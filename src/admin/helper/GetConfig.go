package helper

import (
	"github.com/Unknwon/goconfig"
	"log"
)

var Config *goconfig.ConfigFile
var err error

// 初始化配置文件
func InitConfig() {
	Config, err = goconfig.LoadConfigFile("./config.ini")
	checkConfigErr(err, "获取配置文件失败")
}

// 是否开启 debug 模式
func GetisDebug() string {
	isdebug, _ := Config.GetValue(goconfig.DEFAULT_SECTION, "isDebug")
	return isdebug
}

// 服务器读数据库地址
func GetreadDbServer() string {
	server, _ := Config.GetValue(goconfig.DEFAULT_SECTION, "readDbServer")
	return server
}
// 服务器读数据库端口
func GetreadDbPort() string {
	port, _ := Config.GetValue(goconfig.DEFAULT_SECTION, "readDbPort")
	return port
}
// 服务器读数据库用户名
func GetreadDbUser() string {
	user, _ := Config.GetValue(goconfig.DEFAULT_SECTION, "readDbUser")
	return user
}
// 服务器读数据库密码
func GetreadDbPassword() string {
	password, _ := Config.GetValue(goconfig.DEFAULT_SECTION, "readDbPassword")
	return password
}
// 服务器读数据库数据库
func GetreadDbDatabase() string {
	database, _ := Config.GetValue(goconfig.DEFAULT_SECTION, "readDbDatabase")
	return database
}

// 服务器写数据库地址
func GetwriterDbServer() string {
	server, _ := Config.GetValue(goconfig.DEFAULT_SECTION, "writerDbServer")
	return server
}
// 服务器写数据库端口
func GetwriterDbPort() string {
	port, _ := Config.GetValue(goconfig.DEFAULT_SECTION, "writerDbPort")
	return port
}
// 服务器写数据库用户名
func GetwriterDbUser() string {
	user, _ := Config.GetValue(goconfig.DEFAULT_SECTION, "writerDbUser")
	return user
}
// 服务器写数据库密码
func GetwriterDbPassword() string {
	password, _ := Config.GetValue(goconfig.DEFAULT_SECTION, "writerDbPassword")
	return password
}
// 服务器写数据库数据库
func GetwriterDbDatabase() string {
	database, _ := Config.GetValue(goconfig.DEFAULT_SECTION, "writerDbDatabase")
	return database
}

// sql 文件
func GetsqlLogFile() string {
	file, _ := Config.GetValue(goconfig.DEFAULT_SECTION, "sqlLogFile")
	return file
}
// debug 文件
func GetdebugLogFile() string {
	file, _ := Config.GetValue(goconfig.DEFAULT_SECTION, "debugLogFile")
	return file
}
// error 文件
func GeterrorLogFile() string {
	file, _ := Config.GetValue(goconfig.DEFAULT_SECTION, "errorLogFile")
	return file
}

// 检查是否错误
func checkConfigErr(err error, msg string) {
	if err != nil {
		log.Print("[Log.log.Print][msg]"+msg+"-[err.Error]", err)
		return
	}
}


