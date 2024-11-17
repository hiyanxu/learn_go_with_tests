package main

import (
	"context"
	"fmt"
	"os"
	"time"

	// 导入logrus日志包，别名为log
	log "github.com/sirupsen/logrus"
)

func main() {
	// 设置日志等级
	log.SetLevel(log.WarnLevel)
	// 设置日志输出到什么地方去
	// 将日志输出到标准输出，就是直接在控制台打印出来。
	log.SetOutput(os.Stdout)
	// 设置为true则显示日志在代码什么位置打印的
	//log.SetReportCaller(true)

	// 设置日志以json格式输出， 如果不设置默认以text格式输出
	log.SetFormatter(&log.JSONFormatter{})

	// 打印日志
	log.Debug("调试信息")
	log.Info("提示信息")
	log.Warn("警告信息")
	log.Error("错误信息")
	//log.Panic("致命错误")
	//
	// 为日志加上字段信息，log.Fields其实就是map[string]interface{}类型的别名
	log.WithFields(log.Fields{
		"user_id":    1001,
		"ip":         "192.168.0。100",
		"request_id": "ec2bf8e55a11474392f8867e92624e04",
	}).Info("用户登陆失败.")


	ctx2 := context.Background()
	fmt.Println(ctx2)

	a(ctx2)
	fmt.Println(ctx2)
}

func a(ctx context.Context) {
	ctx3, _ := context.WithTimeout(ctx, 3*time.Second)
	fmt.Println(ctx3)
}
