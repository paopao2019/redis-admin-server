package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"redis-admin-server/model"
	"redis-admin-server/service"
)

func SendCommand(c *gin.Context) {
	var getRedisCommand model.GETRedisCommand
	_ = c.ShouldBindJSON(&getRedisCommand)

	if err, results := service.SendCommand(getRedisCommand); err != nil {
		fmt.Printf("获取失败: %v", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": "获取失败: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg": "获取成功",
			"data": results,   // redis返回的数据
		})

	}

}