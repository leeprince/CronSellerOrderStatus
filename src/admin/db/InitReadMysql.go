package db

import (
	"CronSellerOrderStatus/src/admin/helper"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	//"github.com/go-xorm/core"
	"os"
	"xorm.io/core"
)

var Engine *xorm.Engine
var errEngine error

// 初始化读库
func InitReadMysql() {
	mysqlConfig := helper.GetreadDbUser()+":"+helper.GetreadDbPassword()+"@("+helper.GetreadDbServer()+":"+helper.GetreadDbPort()+")/"+helper.GetwriterDbDatabase()+"?charset=utf8"
	if helper.GetisDebug() == "true" {
		fmt.Println(mysqlConfig)
	}
	helper.DebugLog("InitReadMysql ...", "")
	Engine, errEngine = xorm.NewEngine("mysql", mysqlConfig)
	helper.CheckErr(errEngine, "FailOpenReadMysql.")
	Engine.ShowSQL(true) // 在控制台打印生成 sql 语句
	Engine.SetMaxOpenConns(1000) //最大连接数
	Engine.SetMaxIdleConns(1000) //最大空闲连接数
	Engine.SetTableMapper(core.SnakeMapper{})

	GetsqlLogFile := helper.GetsqlLogFile()
	_, cherr := helper.CheckDirAndCreateDir(GetsqlLogFile)
	helper.CheckErr(cherr, "创建 sql 日志错误")
	f, ferr := os.OpenFile(GetsqlLogFile, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)
	helper.CheckErr(ferr, "CreateSqlLogFail")
	defer f.Close()
	Engine.SetLogger(xorm.NewSimpleLogger(f))
	errEngine = Engine.Ping()
	helper.CheckErr(errEngine, "Engine.Ping() 错误")
	helper.DebugLog("InitReadMysql end ...", "")
}

func CloseDb() {
	helper.DebugLog("CloseReadMysql.ping():", Engine.Ping())
	defer Engine.Close()
}
