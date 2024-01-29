package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"uuid_srv/pkg/snowflake"
)

type UuidReq struct {
	Uuid string `binding:"required" title:"唯一ID" json:"uuid" form:"uuid"`
}

type UuidsReq struct {
	Num int `binding:"required,min=1,max=1000" title:"数量" json:"num" form:"num"`
}
type UuidsRes struct {
	Uuid string `json:"uuid" description:"唯一ID"`
}

// GetUuid 获取唯一ID
func GetUuid(c *gin.Context) {
	var params UuidsReq
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}

	var Uuids = make([]UuidsRes, params.Num)
	for i := 0; i < params.Num; i++ {
		Uuids[i] = UuidsRes{
			Uuid: fmt.Sprintf("%d", snowflake.Instance().NextVal()),
		}
	}

	c.JSON(http.StatusOK, Response{
		SUCCESS,
		Uuids,
		"",
	})
}

// GetDeviceID 获取数据中心ID
func GetDeviceID(c *gin.Context) {
	var params UuidReq
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}

	uuid, err := strconv.ParseInt(params.Uuid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		SUCCESS,
		snowflake.GetDeviceID(uuid),
		"",
	})
}

// GetWorkerId 获取机器ID
func GetWorkerId(c *gin.Context) {
	var params UuidReq
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}

	uuid, err := strconv.ParseInt(params.Uuid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		SUCCESS,
		snowflake.GetWorkerId(uuid),
		"",
	})
}

// GetTimestamp 获取时间戳
func GetTimestamp(c *gin.Context) {
	var params UuidReq
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}

	uuid, err := strconv.ParseInt(params.Uuid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		SUCCESS,
		snowflake.GetTimestamp(uuid),
		"",
	})
}

// GetGenTimestamp 获取创建ID时的时间戳
func GetGenTimestamp(c *gin.Context) {
	var params UuidReq
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}

	uuid, err := strconv.ParseInt(params.Uuid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		SUCCESS,
		snowflake.GetGenTimestamp(uuid),
		"",
	})
}

// GetGenTime 获取创建ID时的时间字符串(精度：秒)
func GetGenTime(c *gin.Context) {
	var params UuidReq
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}

	uuid, err := strconv.ParseInt(params.Uuid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			ERROR,
			map[string]interface{}{},
			err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		SUCCESS,
		snowflake.GetGenTime(uuid),
		"",
	})
}
