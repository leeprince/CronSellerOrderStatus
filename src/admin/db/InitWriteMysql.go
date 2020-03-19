package db

import (
	"CronSellerOrderStatus/src/admin/helper"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	//"github.com/go-xorm/core"

	"xorm.io/core"
	"os"
)

var WriterEngine *xorm.Engine
var errWriterEngine error

// 初始化读库
func InitWriterMysql() {
	mysqlConfig := helper.GetwriterDbUser()+":"+helper.GetwriterDbPassword()+"@("+helper.GetwriterDbServer()+":"+helper.GetwriterDbPort()+")/"+helper.GetwriterDbDatabase()+"?charset=utf8"
	if helper.GetisDebug() == "true" {
		fmt.Println(mysqlConfig)
	}
	fmt.Println("InitReadMysql ...")
	WriterEngine, errWriterEngine = xorm.NewEngine("mysql", mysqlConfig)
	helper.CheckErr(errWriterEngine, "FailOpenReadMysql.")
	WriterEngine.ShowSQL(true) // 在控制台打印生成 sql 语句
	WriterEngine.SetMaxOpenConns(1000) //最大连接数
	WriterEngine.SetMaxIdleConns(1000) //最大空闲连接数
	WriterEngine.SetTableMapper(core.SnakeMapper{})

	f, ferr := os.OpenFile(helper.GetsqlLogFile(), os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)
	helper.CheckErr(ferr, "CreateSqlLogFail")
	defer f.Close()
	errWriterEngine = WriterEngine.Ping()
	helper.CheckErr(errWriterEngine, "WriterEngine.Ping() 错误")
	helper.DebugLog("InitWriterMysql end ...", "")
}

func CloseWriterDb() {
	helper.DebugLog("InitWriterMysql.ping():", WriterEngine.Ping())
	defer Engine.Close()
}
