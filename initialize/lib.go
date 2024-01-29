package initialize

import (
	"fmt"
	"uuid_srv/pkg/snowflake"
)

func LibInit() {
	// 雪花算法
	fmt.Println("雪花算法-例化：start...")
	_, err := snowflake.NewSnowflake(1, 1)
	if err != nil {
		fmt.Println("雪花算法-实例化:Fail-", err)
	}
	fmt.Println("雪花算法-实例化：done")
}
