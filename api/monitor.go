package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"redis-admin-server/model"
)

func GetInfoItemMonitorData(c *gin.Context)  {
	var queryInfo model.RedisMonitorQueryParams
	_ = c.ShouldBindJSON(&queryInfo)
	fmt.Println(queryInfo)
	c.String(http.StatusOK, "ok")
}
