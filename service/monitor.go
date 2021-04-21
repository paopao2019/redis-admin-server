package service

import (
	"fmt"
	"log"
	"redis-admin-server/global"
	"redis-admin-server/model"
	"redis-admin-server/uitls"
	"strconv"
	"strings"
	"time"
)

// 定时监控
func RedisMonitor() {
	var ch chan int
	ticker := time.NewTicker(time.Second * 30)
	go func() {
		for range ticker.C {
			go func() {
				fmt.Println("定时监控:", time.Now().Format("2006-01-02 15:04:05"))
			}()

			// 监控任务
			go getRedisInfo()

		}
		ch <- 1
	}()
	<-ch
}

// 定时采集数据的 函数
func getRedisInfo() {
	var redisClusterList []model.RedisCluster
	err := global.DB.Model(&model.RedisCluster{}).Find(&redisClusterList).Error
	if err != nil {
		log.Println("查询数据库失败", err.Error())
	}
	// 1. 遍历集群
	for _, redisNodes := range redisClusterList {
		var redisClusterInfo model.RedisCluster
		redisClusterInfo.RedisNodeInfo = new(model.RedisNodeInfo)
		// 定义集群状态
		redisClusterInfo.ClusterState = "health"
		redisClusterInfo.ID = redisNodes.ID
		// 1. 遍历集群中的节点
		for _, node := range strings.Split(redisNodes.Nodes, ",") {
			log.Printf("监控节点: %s", node)
			// 1. 定义数据模型
			var redisNodeInfo = new(model.RedisNodeInfo)
			var redisNodeOtherInfo = new(model.RedisNodeOtherInfo)
			var redisMonitorInfo  model.RedisMonitorInfo
			var redisNode = new(model.RedisNode)
			// base信息赋值
			redisMonitorInfo.ClusterID = redisNodes.ID
			redisMonitorInfo.LinkState = "health"
			redisMonitorInfo.Node = node
			//  2. 获取redis原始数据
			err, info := uitls.GetRedisInfo(node, redisNodes.Password)
			if err != nil {
				log.Println("连接redis失败", node, redisNodes.Password,err.Error())
				redisMonitorInfo.LinkState = "bad"
				redisClusterInfo.ClusterState = "bad"
			}

			if redisMonitorInfo.LinkState == "health" {
				// 3. 格式化数据成节点信息需要的
				err = uitls.FormatRedisInfo(info, redisNodeInfo, redisNodeOtherInfo)
				// 赋值采集的监控信息
				redisMonitorInfo.RedisNodeInfo = redisNodeInfo
				redisMonitorInfo.RedisNodeOtherInfo = redisNodeOtherInfo
				redisClusterInfo.RedisNodeInfo = redisNodeInfo
			}

			// 4. 插入监控的的数据库中
			err = global.DB.Create(&redisMonitorInfo).Error
			if err != nil {
				log.Println("插入数据库失败", err.Error())
			}

			// 5. 更新redis节点库 - (当前信息)
			db := global.DB.Where("node = ? ", node).First(&redisNode)
			redisNode.RedisNodeInfo = redisNodeInfo
			redisNode.LinkState =  redisMonitorInfo.LinkState
			err = db.Save(&redisNode).Error
			if err != nil {
				log.Println("更新数据库", err.Error())
			}

		}

		// 2. 更新集群信息到数据库中
		err = global.DB.Where("id = ? ", redisClusterInfo.ID).First(&model.RedisCluster{}).Updates(model.RedisCluster{RedisNodeInfo: redisClusterInfo.RedisNodeInfo, ClusterState: redisClusterInfo.ClusterState}).Error
		if err != nil {
			log.Println("更新数据库", err.Error())
		}

	}

}

// 接口函数
func GetInfoItemMonitorData(queryInfo model.RedisMonitorQueryParams) (err error, resultList []interface{}) {

	for _, node := range queryInfo.NodeList {
		log.Printf("查询节点: %s 的信息", node)
		db := global.DB.Model(&model.RedisMonitorInfo{})

		startTime, _ := strconv.ParseInt(queryInfo.StartTime, 10 ,64)
		endTime, _ := strconv.ParseInt(queryInfo.EndTime, 10 ,64)
		db = db.Debug().Where("updated_at > ? and updated_at < ? and cluster_id = ?",
			time.Unix(startTime,0).Format("2006-01-02 15:04:05"),
			time.Unix(endTime,0).Format("2006-01-02 15:04:05"), queryInfo.ClusterID)
		var nodeMonitorData []model.RedisMonitorInfo
		err := db.Where("node = ? ", node).Find(&nodeMonitorData).Error
		if err != nil {
			log.Println("数据库查询失败,", err.Error())
		}
		resultList = append(resultList, nodeMonitorData)
	}

	return  err, resultList




}