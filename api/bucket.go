package api

import (
	"context"
	"net/http"

	sdktypes "github.com/bnb-chain/greenfield-go-sdk/types"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type bucketInfo struct {
	Name              string `json:"name"`
	ID                string `json:"id"`
	LastUpdateTime    int64  `json:"last_update_at"`
	Owner             string `json:"owners"`
	Status            string `json:"status"`
	ActiveObjectCount uint   `json:"active_object_count"`
}

func (a *API) GetBuckets(c *gin.Context) {
	status, buckets := http.StatusOK, make([]bucketInfo, 0)
	defer func() {
		c.JSON(status, buckets)
	}()

	ctx, cancelCreateBucket := context.WithCancel(context.Background())
	defer cancelCreateBucket()

	spInfo, err := a.cli.ListStorageProviders(ctx, true)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	bucketListRes, err := a.cli.ListBuckets(ctx, sdktypes.ListBucketsOptions{
		ShowRemovedBucket: false,
		Endpoint:          spInfo[0].Endpoint,
	})
	if err != nil {
		log.Printf(err.Error())
		return
	}

	if len(bucketListRes.Buckets) == 0 {
		return
	}

	for _, bucket := range bucketListRes.Buckets {
		buckets = append(buckets, toBucketInfo(bucket))
	}
}

func toBucketInfo(b *sdktypes.BucketMetaWithVGF) bucketInfo {
	return bucketInfo{
		Name:              b.BucketInfo.BucketName,
		ID:                b.BucketInfo.Id.String(),
		LastUpdateTime:    b.UpdateTime,
		Owner:             b.BucketInfo.Owner,
		Status:            b.BucketInfo.BucketStatus.String(),
		ActiveObjectCount: 0, // todo
	}
}
