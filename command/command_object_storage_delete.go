package command

import (
	"fmt"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"strings"
)

func ObjectStorageDelete(ctx Context, params *DeleteObjectStorageParam) error {
	if ctx.NArgs() != 1 {
		return fmt.Errorf("ObjectStorageDelete is failed: %s", "Only two argument can be specified")
	}

	// path
	path := ""
	path = ctx.Args()[0]
	if path != "" && strings.HasPrefix(path, "/") {
		path = strings.Replace(path, "/", "", 0)
	}

	// on SakuraCloud, bucket name is same as AccessKey
	if params.Bucket == "" {
		params.Bucket = params.AccessKey
	}

	auth, err := aws.GetAuth(params.AccessKey, params.SecretKey)
	if err != nil {
		return fmt.Errorf("ObjectStorageDelete is failed: %s", err)
	}
	client := s3.New(auth, aws.Region{
		Name:       "us-west-2",
		S3Endpoint: "https://b.sakurastorage.jp",
	})

	bucket := client.Bucket(params.Bucket)

	// Del
	err = bucket.Del(path)
	if err != nil {
		return fmt.Errorf("ObjectStorageDelete is failed: %s", err)
	}

	return nil
}
