package main

import (
	"github.com/gin-gonic/gin"
	"redis-admin-server/api"
	"redis-admin-server/global"
	"redis-admin-server/initialize"
	"redis-admin-server/middleware"
	"redis-admin-server/service"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Cors())

	// 数据库mysql
	global.DB = initialize.Gorm()

	station := r.Group("/redis")
	{
		station.POST("/addCluster", api.AddRedisCluster)
		station.POST("/getClusterList", api.GetRedisClusterList)
		station.DELETE("/deleteCluster", api.DeleteRedisCluster)
		station.POST("/getClusterById", api.GetRedisClusterById)
		station.PUT("/updateCluster", api.UpdateRedisCluster)
	}

	// 监控数据路由
	r.POST("/monitor/getInfoItemMonitorData", api.GetInfoItemMonitorData)
	// 接受命令操作
	r.POST("/command/sendCommand", api.SendCommand)

	//定时监控任务
	go service.RedisMonitor()

	r.Run(":9996")



}
