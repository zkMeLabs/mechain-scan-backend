package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupInfo struct {
	Owner                  string `json:"owner"`
	Name                   string `json:"group_name"`
	ID                     uint64 `json:"id"`
	LastUpdateAt           int64  `json:"last_update_at"`
	ActiveGroupMemberCount uint64 `json:"active_group_member_count"`
}

func GetGroups(c *gin.Context) {
	groupInfos := make([]GroupInfo, 0)
	c.JSON(http.StatusOK, groupInfos)
}
