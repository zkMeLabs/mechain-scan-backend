package api

import (
	"context"
	"log"
	"net/http"
	"strconv"

	sdktypes "github.com/bnb-chain/greenfield-go-sdk/types"
	"github.com/gin-gonic/gin"
)

type groupInfo struct {
	Owner                  string `json:"owner"`
	Name                   string `json:"group_name"`
	ID                     uint64 `json:"id"`
	LastUpdateAt           int64  `json:"last_update_at"`
	ActiveGroupMemberCount uint64 `json:"active_group_member_count"`
}

func (a *API) GetGroups(c *gin.Context) {
	status, groupInfos := http.StatusOK, make([]groupInfo, 0)
	defer func() {
		c.JSON(status, groupInfos)
	}()
	ctx := context.Background()

	initStartKey := ""
	for {
		groupList, err := a.cli.ListGroupsByOwner(ctx,
			sdktypes.GroupsOwnerPaginationOptions{Limit: maxListMemberNum, StartAfter: initStartKey})
		if err != nil {
			log.Println(err.Error())
			return
		}
		for _, g := range groupList.Groups {
			groupInfos = append(groupInfos, toGroupInfo(g))
		}

		memberNum := len(groupList.Groups)
		if memberNum < maxListMemberNum {
			break
		}

		id := groupList.Groups[memberNum-1].Group.Id
		initStartKey = strconv.FormatUint(id.Uint64(), 10)
	}
}

func toGroupInfo(g *sdktypes.GroupMembers) groupInfo {
	return groupInfo{
		Owner:                  g.Group.Owner,
		Name:                   g.Group.GroupName,
		ID:                     g.Group.Id.Uint64(),
		LastUpdateAt:           g.UpdateTime,
		ActiveGroupMemberCount: 0, // todo
	}
}
