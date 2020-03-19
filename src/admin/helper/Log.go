package helper

import (
	"fmt"
	"log"
	"os"
)

// debug 日志文件
func DebugLog(msg string, data interface{}) {

	GetisDebug := GetisDebug()
	if GetisDebug == "true" {
		fmt.Println("[DebugLog][msg]"+msg+"-[data]", data)

		GetdebugLogFile := GetdebugLogFile()

		_, cherr := CheckDirAndCreateDir(GetdebugLogFile)
		checkErr(cherr, "检查文件并创建文件错误: "+GetdebugLogFile)

		logFile, err := os.OpenFile(GetdebugLogFile, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)
		defer logFile.Close()
		checkErr(err, "打开 Log.debug 日志文件文件错误")
		debugLog := log.New(logFile, "[DEBUG]", log.Lshortfile | log.Ldate | log.Ltime)
		debugLog.Printf("[msg]%v-[data]%v", msg, data)
	}
}

// error 日志文件
func ErrorLog(msg string, data interface{}) {
	GetisDebug := GetisDebug()
	if GetisDebug == "true" {
		log.Print("[ErrorLog][msg]"+msg+"-[data]", data)
	}

	GeterrorLogFile := GeterrorLogFile()

	_, cherr := CheckDirAndCreateDir(GeterrorLogFile)
	checkErr(cherr, "检查文件并创建文件错误"+GeterrorLogFile)
	logFile, err := os.OpenFile(GeterrorLogFile, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)
	defer logFile.Close()
	checkErr(err, "打开 Log.error 日志文件文件错误")
	errorLog := log.New(logFile, "[ERROR]", log.Lshortfile | log.Ldate | log.Ltime)
	errorLog.Printf("[msg]%v-[data]%v", msg, data)
}

// 检查是否错误
func checkErr(err error, msg string) {
	if err != nil {
		log.Print("[Log.log.Print][msg]"+msg+"-[err.Error]", err)
		return
	}
}