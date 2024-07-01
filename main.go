package main

import (
	"uuid_srv/core"
	"uuid_srv/global"
	"uuid_srv/initialize"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper

	// 初始化
	initialize.LibInit()

	// 运行服务
	core.RunWindowsServer()
}
