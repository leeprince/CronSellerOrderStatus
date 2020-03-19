package helper

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// 检查是否错误
func CheckErr(err error, msg string) {
	if err != nil {
		ErrorLog("[CheckErr 检查发生错误]msg:"+msg+"[写入错误日志]err:", err)
		return
	}
}

// int64 转字符串
func StrconvInt64ToString(i int64) string {
	//return strconv.Itoa(int(i)) // 方式一
	return strconv.FormatInt(int64(i), 10) // 方式二
}

// 字符串转 int64
func StrconvStringToInt64(s string) int64 {
	sint64, _ :=  strconv.ParseInt(s, 10, 64)
	return sint64
}

// 字符串转浮点型
func StrconvStringToFloat64(s string) float64 {
	sfloat64, _ :=  strconv.ParseFloat(s, 64)
	return sfloat64
}


// 判断变量的数据类型
func CheckTypeOf(i interface{}) {
	fmt.Printf("变量类型为:%T\n", i)
}

// 判断变量的数据类型
func CheckTypeOfExit(i interface{}) {
	fmt.Printf("变量类型为:%T\n", i)
	os.Exit(0)
}

// 获取商户订单编号: 14 为固定格式时间 + 3 位毫秒数 + 9 位 随机数
func GetOutTradeNo(pre string) string {
	// 14 位时间格式
	ctime := time.Now().Format("20060102150405")

	// 3 位毫秒数
	ntime := time.Now().UnixNano()/ 1e6 // 包含毫秒数的时间戳： 10 位秒级时间戳 + 3 位毫秒数
	stringtime := strconv.Itoa(int(ntime)) // 转化位字符串
	stime := string([]rune(stringtime)[10:13]) // 截取 3 位毫秒数

	// 9 位随机数
	randNum := RandInt64(100000000, 999999999)
	stringNum := StrconvInt64ToString(randNum)

	return pre+ctime+stime+stringNum
}

// 生成区间随机数
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

// 获取当前目录路径
func GetCurrentDirPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))  //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

// 检查文件的路径是否存在， 不存在则创建
func CheckDirAndCreateDir(filePath string) (bool, error) {
	dir := path.Dir(filePath)
	_, statErr := os.Stat(dir)
	if statErr == nil {
		return true, nil // 文件已存在
	}

	mkdirErr := os.MkdirAll(dir, 0755) // 文件不存在且创建多级目录
	if mkdirErr != nil {
		return false, mkdirErr
	}
	return true, nil
}
