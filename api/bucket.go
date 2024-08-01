package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BucketInfo struct {
	// owner is the account address of bucket creator, it is also the bucket owner.
	Owner string `json:"owners"`
	// bucket_name is a globally unique name of bucket
	Name string `json:"name"`
	// id is the unique identification for bucket.
	ID uint `json:"id"`
	// create_at define the block timestamp when the bucket created.
	CreateAt          int64  `json:"create_at"`
	Status            string `json:"status"`
	ActiveObjectCount uint   `json:"active_object_count"`
}

func GetBuckets(c *gin.Context) {
	// c, cancelCreateBucket := context.WithCancel(c.Request.Context())
	// defer cancelCreateBucket()

	// spInfo, err := client.ListStorageProviders(c, true)
	// if err != nil {
	// 	fmt.Println("fail to get SP info to list bucket:", err.Error())
	// 	return nil
	// }

	// bucketListRes, err := client.ListBuckets(c, sdktypes.ListBucketsOptions{
	// 	ShowRemovedBucket: false,
	// 	Endpoint:          spInfo[0].Endpoint,
	// })
	// if err != nil {
	// 	return err
	// }

	// if len(bucketListRes.Buckets) == 0 {
	// 	return nil
	// }

	// for _, bucket := range bucketListRes.Buckets {
	// 	info := bucket.BucketInfo

	// 	location, _ := time.LoadLocation("Asia/Shanghai")
	// 	t := time.Unix(info.CreateAt, 0).In(location)
	// 	if !bucket.Removed {
	// 		fmt.Printf("%s  %s\n", t.Format(iso8601DateFormat), info.BucketName)
	// 	}
	// }
	buckets := make([]BucketInfo, 0)
	c.JSON(http.StatusOK, buckets)
}
