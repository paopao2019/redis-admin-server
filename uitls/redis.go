package uitls

import (
	"github.com/go-redis/redis"
	"redis-admin-server/model"
	"strconv"
	"strings"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func GetRedisInfo(node string, password string ) (err error, info string) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     node,
		Password: password,
		DB:       0,  // use default DB
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err, info
	}


	info, err =  rdb.Info().Result()
	//fmt.Println("-------", info)
	return nil, info
}




func FormatRedisInfo(info string, redisNodeInfo *model.RedisNodeInfo, redisNodeOtherInfo *model.RedisNodeOtherInfo) (err error) {
	var dbKeysArray []int
	infoArray := strings.Split(strings.Trim(info, "\n"), "\n")
	for _ ,v := range infoArray {
		// 统一个里面的所有元素都 除去 \r 字符
		v = strings.Trim(v, "\r")
		if strings.HasPrefix(v, "uptime_in_seconds") {
			redisNodeInfo.UptimeInSeconds = strings.Split(v, ":")[1]
		}
		if strings.HasPrefix(v, "redis_version") {
			redisNodeInfo.RedisVersion = strings.Split(v, ":")[1]
		}
		if strings.HasPrefix(v, "redis_mode") {
			redisNodeInfo.RedisMode = strings.Split(v, ":")[1]
		}
		if strings.HasPrefix(v, "os") {
			redisNodeInfo.OS = strings.Split(v, ":")[1]
		}
		if strings.HasPrefix(v, "used_memory:") {
			redisNodeInfo.UsedMemory = strings.Split(v, ":")[1]
		}
		if strings.HasPrefix(v, "total_system_memory_human") {
			redisNodeInfo.TotalSystemMemoryHuman = strings.Split(v, ":")[1]
		}
		if strings.HasPrefix(v, "role") {
			redisNodeInfo.NodeRole = strings.Split(v, ":")[1]
		}
		if strings.HasPrefix(v, "db") {
			dbKeys, _ :=strconv.Atoi(strings.Split(strings.Split(strings.Split(v, ":")[1], ",")[0], "=")[1])
			dbKeysArray = append(dbKeysArray, dbKeys)
		}
		// 所有db 中的 keys总数求和
		redisNodeInfo.TotalKeys = Sum(dbKeysArray)
		redisNodeInfo.DBSize = len(dbKeysArray)


		// redisNodeOtherInfo 信息
		if strings.HasPrefix(v, "used_memory_peak:") {
			redisNodeOtherInfo.UsedMemoryPeak = strings.Split(v, ":")[1]
		}
		if strings.HasPrefix(v, "used_memory_rss:") {
			redisNodeOtherInfo.UsedMemoryRss = strings.Split(v, ":")[1]
		}
		if strings.HasPrefix(v, "connected_clients") {
			redisNodeOtherInfo.ConnectedClients = strings.Split(v, ":")[1]
		}
		if strings.HasPrefix(v, "blocked_clients") {
			redisNodeOtherInfo.BlockedClients = strings.Split(v, ":")[1]
		}
		if strings.HasPrefix(v, "expired_keys") {
			redisNodeOtherInfo.ExpiredKeys = strings.Split(v, ":")[1]
		}
		if strings.HasPrefix(v, "total_commands_processed") {
			redisNodeOtherInfo.TotalCommandsProcessed = strings.Split(v, ":")[1]
		}
		if strings.HasPrefix(v, "keyspace_hits") {
			redisNodeOtherInfo.KeySpaceHits = strings.Split(v, ":")[1]
		}
		if strings.HasPrefix(v, "keyspace_misses") {
			redisNodeOtherInfo.KeySpaceMisses = strings.Split(v, ":")[1]
		}
	}
	//fmt.Println("----------------",redisNodeInfo.NodeRole)
	//fmt.Println("----------------",redisNodeInfo.TotalSystemMemoryHuman)
	//fmt.Println("----------------",redisNodeInfo.UsedMemoryHuman)
	//fmt.Println("----------------",redisNodeInfo.OS)
	//fmt.Println("----------------",redisNodeInfo.TotalKeys)
	//fmt.Println("----------------",redisNodeInfo.DBSize)
	return  err
}