package model


type RedisMonitorQueryParams struct {
	ClusterID       uint `json:"cluster_id"`
	StartTime string `json:"start_time"`   // 时间戳格式 s 如: 1618472090
	EndTime string `json:"end_time"`       // 时间戳格式 s 如: 1618472090
	NodeList []string `json:"node_list"`
}
