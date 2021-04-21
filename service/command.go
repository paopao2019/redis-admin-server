package service

import (
	"github.com/go-redis/redis"
	"log"
	"redis-admin-server/global"
	"redis-admin-server/model"
	"strings"
)

func SendCommand(command model.GETRedisCommand) (err error, result string) {

	// 1. 获取命令 获取redis中的集群 master和密码
	var redisCluster model.RedisCluster
	err = global.DB.Where("id = ?", command.ClusterID).First(&redisCluster).Error
	if err != nil {
		log.Println("数据库链接失败", err.Error())
	}

	// 2. 连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisCluster.ClusterMaster,
		Password: redisCluster.Password,
		DB:       command.Database,  // use default DB
	})
	_, err = rdb.Ping().Result()

	if err != nil {
		return err, result
	}

	if command.Command != "" {
		commandArray := strings.Fields(command.Command)
		//  []string 转成 []interface{}
		var commandInterface []interface{}
		for _, value := range commandArray {
			commandInterface = append(commandInterface, value)
		}

		// 3. 得到数据
		if strings.Contains(command.Command, "keys") {
			resultArray, _ := rdb.Keys(commandArray[1]).Result()
			result = strings.Join(resultArray, "\n")
		} else {
			// Db 不支持 keys功能
			result, err = rdb.Do(commandInterface...).String()
		}

		// 命令错误 或者key不存在错误  都不做为错误 而应该做为返回结果
		if err != nil && (err.Error() == "redis: nil" || strings.Contains(err.Error(), "ERR")){
			result = err.Error()
			err = nil
		}
		log.Println(result)
	}

	return err, result
}


