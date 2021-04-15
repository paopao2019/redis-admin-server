package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"redis-admin-server/model"
	"redis-admin-server/service"
)

// 增
func AddRedisCluster(c *gin.Context) {
	var redisCluster model.RedisCluster
	_ = c.ShouldBindJSON(&redisCluster)
	if err := service.AddRedisCluster(redisCluster); err != nil {
		fmt.Printf("创建失败: %v", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": "创建失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg": "创建成功",
		})
	}
}


// 删
func DeleteRedisCluster(c *gin.Context) {
	var redisCluster model.GetById
	_ = c.ShouldBindJSON(&redisCluster)

	if err := service.DeleteRedisCluster(redisCluster.Id); err != nil {
		fmt.Printf("创建失败: %v", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": "删除失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg": "删除成功",
		})
	}
}

// 改
func UpdateRedisCluster(c *gin.Context) {
	var R model.RedisCluster
	_ = c.ShouldBindJSON(&R)

	if err := service.UpdateRedisCluster(R); err != nil {
		fmt.Printf("更新失败: %v", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": "更新失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg": "更新成功",
		})
	}
}

// 查通过Id
func GetRedisClusterById(c *gin.Context) {
	var reqId model.GetById
	_ = c.ShouldBindJSON(&reqId)

	if err, results := service.GetRedisClusterById(reqId.Id); err != nil {
		fmt.Printf("获取失败: %v", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": "获取失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg": "获取成功",
			"data": map[string]interface{} {
				"redis": results,
			},
		})

	}

}

// 查
func GetRedisClusterList(c *gin.Context) {
	var pageInfo model.SearchRedisClusterParams
	_ = c.ShouldBind(&pageInfo)

	if err, list, total := service.GetRedisClusterList(pageInfo.RedisCluster, pageInfo.PageInfo, pageInfo.Desc); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": "获取失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg": "获取成功",
			"data": map[string]interface{} {
				"list": list,
				"total": total,
				"page": pageInfo.Page,
				"pageSize": pageInfo.PageSize,
			},
		})
	}
}

