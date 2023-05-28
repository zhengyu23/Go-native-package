package main

import (
	"io"
	"log"
	"os"
)

/*
	Go原生包的 Logger是不分级的，但实际应用中我们又需要对 Logger进行诸如
	“追踪、重要信息、警告、错误”等日志等级进行分级，所以根据log原生包自定义
	了分级 Logger。
*/

var (
	Trace   *log.Logger // 追踪级别，不打印任何东西
	Info    *log.Logger // 重要信息级别
	Warning *log.Logger // 警告级别
	Error   *log.Logger // 错误级别
)

func init() {
	file, err := os.OpenFile("errors.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("无法打开 errors.log 文件: ", err)
	}

	Trace = log.New(io.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stdout),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Trace.Println("追踪")
	Info.Println("重要信息")
	Warning.Println("警告")
	Error.Println("错误")
}
