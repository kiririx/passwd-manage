package main

import (
	"github.com/kiririx/passwd-manage/conf"
	"github.com/kiririx/passwd-manage/router"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 设置端口号启动
	router.SetupRouter(conf.Ginner)
	if err := conf.Ginner.Run(":8080"); err != nil {
		panic(err)
	}
}
