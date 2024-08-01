package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ObjectInfo struct {
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

func GetObjects(c *gin.Context) {
	groups := make([]GroupInfo, 0)
	c.JSON(http.StatusOK, groups)
}
