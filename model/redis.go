package model

import (
	"gorm.io/gorm"
)

type RedisCluster struct {
	gorm.Model
	Name   string `json:"name" form:"name" gorm:"column:name;type:varchar(128);not null;comment:集群名称"`
	Nodes    string `json:"nodes" form:"nodes"  gorm:"column:nodes;type:varchar(256);not null;comment:节点 逗号分隔"`
	Mode    string `json:"mode,omitempty" form:"mode"  gorm:"column:mode;type:varchar(1);not null;comment:redis类型 1.single  2.主从master-salve  3.哨兵 sentinel"`
	Status          *bool          `json:"status" form:"status" gorm:"comment:集群状态"`
	Env     string `json:"env" form:"env" gorm:"column:env;type:varchar(128);not null;comment:环境"`
	Password     string `json:"password" form:"env" gorm:"column:password;type:varchar(128);comment:密码"`
	ClusterState string `json:"cluster_state"  gorm:"column:cluster_state;type:varchar(50)"`
	ClusterKnownNodes int `json:"cluster_known_nodes" gorm:"column:cluster_known_nodes;type:int(8)"`
	ClusterMaster string `json:"cluster_master"  gorm:"column:cluster_master;type:varchar(50)"`
	*RedisNodeInfo  // 能连接
}

type SearchRedisClusterParams struct {
	RedisCluster
	PageInfo
	Desc     bool   `json:"desc"`
}


// 需要存入数据库节点中的采集的Redis 信息
type RedisNodeInfo struct {
	RedisVersion string `json:"redis_version" gorm:"column:redis_version;type:varchar(50);comment:版本"`
	RedisMode string `json:"redis_mode" gorm:"column:redis_mode;type:varchar(50);comment:模式"`
	OS string `json:"os" gorm:"column:os;type:varchar(50);comment:操作系统"`
	UptimeInSeconds string `json: uptime_in_seconds gorm: "column: uptime_in_seconds";type:varchar(50)`
	TotalSystemMemoryHuman string  `json: total_system_memory_human gorm: "column: total_system_memory_human";type:varchar(50)`
	UsedMemoryHuman string `json: used_memory_human gorm: "column: used_memory_human";type:varchar(50)`
	DBSize int  `json: db_size gorm: "column: db_size";type:int(8)`
	NodeRole    string `json:"node_role" gorm:"column:node_role;type:varchar(50);comment:节点角色"`
	TotalKeys int `json: total_keys gorm: "column: total_keys";type:int(8)`
}

type RedisNodeOtherInfo struct {
	// 需要监控的其他信息
}

type RedisNode struct {
	gorm.Model
	ClusterID uint `json:"cluster_id" gorm:"column:cluster_id;not null;comment:集群ID"`
	Node    string `json:"node" gorm:"column:node;type:varchar(50);comment:节点"`
	LinkState string `json:"link_state" gorm:"column:link_state;type:varchar(50);comment:redis状态"`
	*RedisNodeInfo  // 能连接 LinkState是好的 才有用
	RedisCluster  RedisCluster `json:"redis_cluster" gorm:"foreignKey:cluster_id;references:id"`

}

// 定时监控的表 监控信息
type RedisMonitorInfo struct {
	RedisNode
	*RedisNodeOtherInfo
}
