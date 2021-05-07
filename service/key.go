package service

import (
	"github.com/go-redis/redis"
	"log"
	"redis-admin-server/global"
	"redis-admin-server/model"
	"redis-admin-server/uitls"
)

func KeyQuery(queryInfo model.KeyScanParams) (keyInfo map[string]interface{},err error) {
	// 1. 获取命令 获取redis中的集群 master和密码
	var redisCluster model.RedisCluster
	err = global.DB.Where("id = ?", queryInfo.ClusterID).First(&redisCluster).Error
	if err != nil {
		log.Println("数据库链接失败", err.Error())
	}

	// 2. 连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisCluster.ClusterMaster,
		Password: redisCluster.Password,
		DB:       queryInfo.Database,  // use default DB
	})

	keyInfo = make(map[string]interface{})
	result, err := rdb.Get(queryInfo.Key).Result()
	keyInfo["value"] = result
	s, err := rdb.Type(queryInfo.Key).Result()
	keyInfo["type"] = s
	duration, err := rdb.TTL(queryInfo.Key).Result()
	keyInfo["ttl"] = duration / 1000/1000/1000
	println(keyInfo , "----------")
	return  keyInfo, err
}

// redis scan key
func KeyScan(queryInfo model.KeyScanParams) (list []string, err error ) {
	// 1. 获取命令 获取redis中的集群 master和密码
	var redisCluster model.RedisCluster
	err = global.DB.Where("id = ?", queryInfo.ClusterID).First(&redisCluster).Error
	if err != nil {
		log.Println("数据库链接失败", err.Error())
	}

	// 2. 连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisCluster.ClusterMaster,
		Password: redisCluster.Password,
		DB:       queryInfo.Database,  // use default DB
	})
	var cursor uint64
	list, cursor, err = rdb.Scan(cursor, queryInfo.Key, queryInfo.Count).Result()

	return list, err
}



//
func GetDBList(cluster_id string) (list []map[string]interface{}, err error) {
	// 1. 获取命令 获取redis中的集群 master和密码
	var redisCluster model.RedisCluster
	err = global.DB.Where("id = ?", cluster_id).First(&redisCluster).Error
	if err != nil {
		log.Println("数据库链接失败", err.Error())
	}

	// 2. 连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisCluster.ClusterMaster,
		Password: redisCluster.Password,
		DB:       0,  // use default DB
	})

	list, err = uitls.GetDBList(rdb)

	return list, err
}