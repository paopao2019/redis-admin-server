package model


// scan 或者 query 都是这个model
type KeyScanParams struct {
	ClusterID       uint `json:"cluster_id"`
	Count int64 `json:"count"`   //
	Database int `json:"database"`       //
	Key string `json:"key"`
}

