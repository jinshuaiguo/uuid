package initialize

import (
	"fmt"
	"uuid_srv/global"
	"uuid_srv/pkg/snowflake"
)

func LibInit() {
	// 雪花算法
	fmt.Println("雪花算法-例化：start...")
	fmt.Println(global.GVA_CONFIG.Snowflake.Workerid)
	_, err := snowflake.NewSnowflake(1, global.GVA_CONFIG.Snowflake.Workerid)
	if err != nil {
		fmt.Println("雪花算法-实例化:Fail-", err)
	}
	fmt.Println("雪花算法-实例化：done")
}
