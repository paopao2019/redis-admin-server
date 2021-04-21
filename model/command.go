package model

type GETRedisCommand struct {
	ClusterID uint `json:"cluster_id" gorm:"column:cluster_id;not null;comment:集群ID"`
	Database    int `json:"database" gorm:"column:database;type:int(8);comment:选择的库"`
	Command string `json:"command" gorm:"column:command;type:varchar(50);comment:redis命令"`
}
