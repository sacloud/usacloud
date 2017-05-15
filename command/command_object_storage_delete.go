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

	// is directory?
	key, _ := bucket.GetKey(path) // if path is directory , GetKey(path) returns nil(with 404 err)
	if key == nil {
		if !params.Recursive {
			return fmt.Errorf("%q is directory. Use -r or --recursive flag", path)
		}

		// path is directory
		if path != "" && !strings.HasSuffix(path, "/") {
			path = path + "/"
		}
		return objectStorageDelRecursive(path, bucket)
	}
	// path is file
	err = objectStorageDel(path, bucket)
	if err != nil {
		return fmt.Errorf("ObjectStorageDelete is failed: %s", err)
	}

	return nil
}

func objectStorageDelRecursive(path string, bucket *s3.Bucket) error {

	res, err := bucket.List(path, "/", "", 0)
	if err != nil {
		return err
	}

	// first, delete each files
	for _, content := range res.Contents {
		err := objectStorageDel(content.Key, bucket)
		if err != nil {
			return err
		}
	}

	// next, delete each dir
	for _, pref := range res.CommonPrefixes {
		err := objectStorageDelRecursive(pref, bucket)
		if err != nil {
			return nil
		}
	}

	return nil
}

func objectStorageDel(path string, bucket *s3.Bucket) error {
	err := bucket.Del(path)
	fmt.Fprintf(GlobalOption.Progress, "Deleted: %s\n", path)
	return err
}
