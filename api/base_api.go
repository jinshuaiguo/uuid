package api

import (
	"github.com/gin-gonic/gin"
)

const (
	ERROR   = 0
	SUCCESS = 1
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Routers(Router *gin.Engine) {
	PublicGroup := Router.Group("")
	{
		PublicGroup.GET("/getUuid", GetUuid)
		PublicGroup.GET("/getDeviceID", GetDeviceID)
		PublicGroup.GET("/getWorkerId", GetWorkerId)
		PublicGroup.GET("/getTimestamp", GetTimestamp)
		PublicGroup.GET("/getGenTimestamp", GetGenTimestamp)
		PublicGroup.GET("/getGenTime", GetGenTime)
	}
}
