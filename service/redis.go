package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"redis-admin-server/global"
	"redis-admin-server/model"
	"redis-admin-server/uitls"
	"strings"
)

func AddRedisCluster(redisCluster model.RedisCluster) (err error) {
	if !errors.Is(global.DB.Where("name = ? ", redisCluster.Name).First(&model.RedisCluster{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同集群")
	}
	err = global.DB.Create(&redisCluster).Error
	if err != nil {
		return err
	}
	// 1. 获取创建node表  其他信息 通过 redis-api获取
	redisPassword := redisCluster.Password
	redisClusterAddress := redisCluster.Nodes
	redisNodeArray := strings.Split(redisClusterAddress, ",")
	redisCluster.ClusterState = "health"
	redisCluster.ClusterKnownNodes = len(redisNodeArray)
	for _, redisNodeAddr := range redisNodeArray {
		// 获取节点信息
		fmt.Println(redisNodeAddr)
		var redisNodeInfo = new(model.RedisNodeInfo)
		var redisNode model.RedisNode
		redisNode.RedisNodeInfo = redisNodeInfo
		redisNode.Node = redisNodeAddr
		//  1. 获取redis原始数据
		err, info := uitls.GetRedisInfo(redisNodeAddr, redisPassword)
		if err != nil {
			log.Println("连接redis失败", err.Error())
			redisNode.LinkState = "bad"
			redisCluster.ClusterState = "bad"
		} else {
			redisNode.LinkState = "health"
		}

		if redisNode.LinkState == "health" {
			// 2. 格式化数据成节点信息需要的
			err = uitls.FormatRedisInfo(info, redisNodeInfo)
			fmt.Printf("redis节点测试: %s\n", redisNodeInfo.NodeRole)
			redisNode.RedisNodeInfo = redisNodeInfo
		}
		redisNode.ClusterID = redisCluster.ID
		// 3. 数据写入 node节点DB中
		err = global.DB.Create(&redisNode).Error
		if err != nil {
			return err
			log.Println("host写入数据库失败", err.Error())
		}

		// 4. 综合节点数据 - 节点中master的信息 在写入cluster中 (集群中的master)
		fmt.Println("redisNode.NodeRole:", redisNode.NodeRole)
		if redisNode.NodeRole == "master" {
			redisCluster.ClusterMaster = redisNodeAddr
			redisCluster.RedisNodeInfo = redisNodeInfo
			//global.DB.Debug().Where("name = ? ", redisCluster.Name).First(&model.RedisCluster{}).Updates(model.RedisCluster{RedisNodeInfo: redisNodeInfo})
			if err != nil {
				log.Println("更新集群中的 信息失败", err.Error())
			}
		} else {
			fmt.Printf("info:%s slave角色的节点", redisNode.Node)
		}
		global.DB.Save(redisCluster)
	}
	return err
}

// 删
func DeleteRedisCluster(id float64) (err error) {
	// 1. 判断不存在没有在其他 表中使用 虚拟机中 <- 待补充
	var redisCluster model.RedisCluster
	// 2. 删除
	err = global.DB.Where("id = ?", id).Delete(&redisCluster).Error
	return err
}

// 改
func UpdateRedisCluster(redisCluster model.RedisCluster) (err error) {
	// 通过唯一ID进行更新
	err = global.DB.Where("id = ?", redisCluster.ID).First(&model.RedisCluster{}).Updates(&redisCluster).Error
	return err
}

// 查GetRedisClusterById
// 查
func GetRedisClusterById(id float64) (err error, redisCluster model.RedisCluster) {
	err = global.DB.Where("id = ?", id).First(&redisCluster).Error
	return err, redisCluster
}


// 查
func GetRedisClusterList(redisCluster model.RedisCluster, info model.PageInfo, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&model.RedisCluster{})
	var redisClusterList []model.RedisCluster
	if redisCluster.Mode != "" {
		db = db.Where("mode = ?", redisCluster.Mode)
	}
	if redisCluster.Name != "" {
		db = db.Where("name like ?", "%" + redisCluster.Name + "%")
	}
	err = db.Count(&total).Error

	if err != nil {
		return err, redisClusterList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		err = db.Find(&redisClusterList).Error
	}
	return err, redisClusterList, total

}
