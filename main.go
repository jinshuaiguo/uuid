package main

import (
	"uuid_srv/core"
	"uuid_srv/initialize"
)

func main() {
	// 初始化
	initialize.LibInit()
	// 运行服务
	core.RunWindowsServer()
}
