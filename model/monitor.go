package model

import "time"


type RedisMonitorQueryParams struct {
	ClusterID       uint `json:"cluster_id"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	NodeList []string `json:"node_list"`
}
