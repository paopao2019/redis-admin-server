package api

import (
"fmt"
"github.com/gin-gonic/gin"
"log"
"net/http"
"redis-admin-server/model"
"redis-admin-server/service"
)

func KeyQuery(c *gin.Context)  {
	var queryInfo model.KeyScanParams
	_ = c.ShouldBindJSON(&queryInfo)
	fmt.Println(queryInfo)  //

	// 查询数据库
	keyInfo ,err := service.KeyQuery(queryInfo)
	if err != nil {
		log.Println("数据库查询失败", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg": "获取成功",
		"data": keyInfo,
	})
}


func KeyScan(c *gin.Context)  {
	var queryInfo model.KeyScanParams
	_ = c.ShouldBindJSON(&queryInfo)
	fmt.Println(queryInfo)  //

	// 查询数据库
	list, err := service.KeyScan(queryInfo)
	if err != nil {
		log.Println("scan失败", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": "scan失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg": "获取成功",
		"data": list,
	})
}


// GetDBList
func GetDBList (c *gin.Context) {
	cluster_id := c.Param("cluster_id")
	if list, err := service.GetDBList(cluster_id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": "获取失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg": "获取成功",
			"data": list,
		})
	}
}