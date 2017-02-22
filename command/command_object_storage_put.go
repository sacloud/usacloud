package command

import (
	"fmt"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"os"
	"path/filepath"
	"strings"
)

func ObjectStoragePut(ctx Context, params *PutObjectStorageParam) error {
	if ctx.NArgs() != 2 {
		return fmt.Errorf("ObjectStoragePut is failed: %s", "Only two argument can be specified")
	}

	// validate filepath
	filePath := ""
	filePath = ctx.Args()[0]
	_, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("file[%s] is not exists: %s", filePath, err)
	}
	// target file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("ObjectStoragePut is failed: %s", err)
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		return fmt.Errorf("ObjectStoragePut is failed: %s", err)
	}

	// path
	path := ""
	path = ctx.Args()[1]
	if path != "" && strings.HasPrefix(path, "/") {
		path = strings.Replace(path, "/", "", 1)
	}

	// if path is dir, set filename from filePath
	if strings.HasSuffix(path, "/") {
		path = fmt.Sprintf("%s%s", path, filepath.Base(filePath))
	}

	// on SakuraCloud, bucket name is same as AccessKey
	if params.Bucket == "" {
		params.Bucket = params.AccessKey
	}

	auth, err := aws.GetAuth(params.AccessKey, params.SecretKey)
	if err != nil {
		return fmt.Errorf("ObjectStoragePut is failed: %s", err)
	}
	client := s3.New(auth, aws.Region{
		Name:       "us-west-2",
		S3Endpoint: "https://b.sakurastorage.jp",
	})

	bucket := client.Bucket(params.Bucket)

	// put
	err = bucket.PutReader(path, file, fi.Size(), params.ContentType, s3.PublicRead)
	if err != nil {
		return fmt.Errorf("ObjectStoragePut is failed: %s", err)
	}

	return nil
}
