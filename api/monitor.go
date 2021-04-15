package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"redis-admin-server/model"
	"redis-admin-server/service"
)

func GetInfoItemMonitorData(c *gin.Context)  {
	var queryInfo model.RedisMonitorQueryParams
	_ = c.ShouldBindJSON(&queryInfo)
	fmt.Println(queryInfo)  //

	// 查询数据库
	err, list := service.GetInfoItemMonitorData(queryInfo)
	if err != nil {
		log.Println("数据库查询失败", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"data": list,
	})
}
