package core

import (
	"fmt"
	"time"
	"uuid_srv/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	// TODO 端口号
	address := fmt.Sprintf(":%d", 8809)

	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	fmt.Printf("server run success on %s\n", address)

	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("start err:%s\n", err.Error())
	}
}
