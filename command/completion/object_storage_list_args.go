package completion

import (
	"fmt"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"path"
	"strings"
)

func ObjectStorageListCompleteArgs(ctx command.Context, params *params.ListObjectStorageParam, cur, prev, commandName string) {

	if ctx.NArgs() != 0 {
		return
	}
	remotePath := cur

	if remotePath != "" && strings.HasPrefix(remotePath, "/") {
		remotePath = strings.Replace(remotePath, "/", "", 1)
	}

	if remotePath != "" && !strings.HasSuffix(remotePath, "/") {
		// search under current dir
		remotePath = path.Dir(remotePath)
		if remotePath == "." {
			remotePath = ""
		} else {
			remotePath += "/"
		}

	}

	// on SakuraCloud, bucket name is same as AccessKey
	if params.Bucket == "" {
		params.Bucket = params.AccessKey
	}

	auth, err := aws.GetAuth(params.AccessKey, params.SecretKey)
	if err != nil {
		return
	}
	client := s3.New(auth, aws.Region{
		Name:       "us-west-2",
		S3Endpoint: "https://b.sakurastorage.jp",
	})

	bucket := client.Bucket(params.Bucket)

	// list all
	res, err := bucket.List(remotePath, "/", "", 0)
	if err != nil {
		return
	}

	// we search current & current+1 depth
	for _, dir := range res.CommonPrefixes {
		if cur != "" {
			unders, err := bucket.List(dir, "/", "", 0)
			if err == nil {
				for _, v := range unders.CommonPrefixes {
					fmt.Println(v)
				}
			}
		}
		fmt.Println(dir)
	}

}
