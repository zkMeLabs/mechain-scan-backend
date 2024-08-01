package api

import (
	"log"
	"net/http"

	sdktypes "github.com/bnb-chain/greenfield-go-sdk/types"
	"github.com/gin-gonic/gin"
)

type objectInfo struct {
	// creator is the address of the uploader, it always be same as owner address
	Creator string `json:"creator,omitempty"`
	// object_name is the name of object
	ObjectName string `json:"object_name,omitempty"`
	// content_type define the format of the object which should be a standard MIME type.
	ContentType string `json:"content_type,omitempty"`
	// payloadSize is the total size of the object payload
	PayloadSize uint64 `json:"payload_size,omitempty"`
	// object_status define the upload status of the object.
	Status string `json:"object_status,omitempty"`
	// visibility defines the highest permissions for object. When an object is public, everyone can access it.
	Visibility string `json:"visibility,omitempty"`
	UpdatedAt  int64  `json:"updated_at,omitempty"`
	// bucket_name is the name of the bucket
	BucketName string `json:"bucket_name,omitempty"`
}

func (a *API) GetObjects(c *gin.Context) {
	status, objects := http.StatusOK, make([]objectInfo, 0)
	defer func() {
		c.JSON(status, objects)
	}()
	bucketName := "mechain" // todo
	var continuationToken string
	for {
		listResult, err := a.cli.ListObjects(c, bucketName, sdktypes.ListObjectsOptions{
			ShowRemovedObject: false,
			MaxKeys:           defaultMaxKey,
			ContinuationToken: continuationToken,
			Prefix:            "",
		})
		if err != nil {
			log.Println(err.Error())
		}

		for _, o := range listResult.Objects {
			objects = append(objects, toObjectInfo(o))
		}
		if !listResult.IsTruncated {
			break
		}
		continuationToken = listResult.NextContinuationToken
	}
}

func toObjectInfo(obj *sdktypes.ObjectMeta) objectInfo {
	return objectInfo{
		Creator:     obj.ObjectInfo.Creator,
		ObjectName:  obj.ObjectInfo.ObjectName,
		ContentType: obj.ObjectInfo.ContentType,
		PayloadSize: obj.ObjectInfo.PayloadSize,
		Status:      obj.ObjectInfo.ObjectStatus.String(),
		Visibility:  obj.ObjectInfo.Visibility.String(),
		UpdatedAt:   obj.ObjectInfo.UpdatedAt,
		BucketName:  obj.ObjectInfo.BucketName,
	}
}
