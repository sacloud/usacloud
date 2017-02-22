package command

import (
	"fmt"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"io"
	"os"
	"strings"
)

func ObjectStorageGet(ctx Context, params *GetObjectStorageParam) error {
	if ctx.NArgs() > 2 {
		return fmt.Errorf("ObjectStorageGet is failed: %s", "Only two argument can be specified")
	}

	// validate filepath
	filePath := ""
	if ctx.NArgs() > 1 {
		filePath = ctx.Args()[1]
	}

	// path
	path := ""
	path = ctx.Args()[0]
	if path != "" && strings.HasPrefix(path, "/") {
		path = strings.Replace(path, "/", "", 1)
	}
	// if path is dir, set filename from filePath
	if strings.HasSuffix(path, "/") {
		return fmt.Errorf("path must not be directory: %s", path)
	}

	// on SakuraCloud, bucket name is same as AccessKey
	if params.Bucket == "" {
		params.Bucket = params.AccessKey
	}

	auth, err := aws.GetAuth(params.AccessKey, params.SecretKey)
	if err != nil {
		return fmt.Errorf("ObjectStorageGet is failed: %s", err)
	}
	client := s3.New(auth, aws.Region{
		Name:       "us-west-2",
		S3Endpoint: "https://b.sakurastorage.jp",
	})

	bucket := client.Bucket(params.Bucket)

	// get key
	data, err := bucket.Get(path)
	if err != nil {
		return fmt.Errorf("ObjectStorageGet is failed: %s", err)
	}

	// write
	var w io.Writer
	if filePath == "" {
		w = GlobalOption.Out
	} else {
		f, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("ObjectStorageGet is failed: %s", err)
		}
		defer f.Close()
		w = f
	}

	_, err = w.Write(data)

	return err
}
