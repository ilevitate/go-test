package main

import (
	"go-test/src/global"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile)
	var err error
	//加载配置
	err = global.LoadConfig()
	if err != nil {
		println("加载配置文件失败：" + err.Error())
		return
	}
	//配置日志
	err = global.SetLogger()
	if err != nil {
		println("设置日志输出失败：" + err.Error())
	}

	//设置时间时区
	global.TimeLocal, err = time.LoadLocation("Local")
	if err != nil {
		global.Logger.Fatal("设置时区失败：" + err.Error())
		return
	}

}
