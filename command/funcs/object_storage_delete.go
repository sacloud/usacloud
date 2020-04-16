// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package funcs

import (
	"fmt"
	"strings"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ObjectStorageDelete(ctx command.Context, params *params.DeleteObjectStorageParam) error {
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
		return objectStorageDelRecursive(ctx, path, bucket)
	}
	// path is file
	err = objectStorageDel(ctx, path, bucket)
	if err != nil {
		return fmt.Errorf("ObjectStorageDelete is failed: %s", err)
	}

	return nil
}

func objectStorageDelRecursive(ctx command.Context, path string, bucket *s3.Bucket) error {

	res, err := bucket.List(path, "/", "", 0)
	if err != nil {
		return err
	}

	// first, delete each files
	for _, content := range res.Contents {
		err := objectStorageDel(ctx, content.Key, bucket)
		if err != nil {
			return err
		}
	}

	// next, delete each dir
	for _, pref := range res.CommonPrefixes {
		err := objectStorageDelRecursive(ctx, pref, bucket)
		if err != nil {
			return nil
		}
	}

	return nil
}

func objectStorageDel(ctx command.Context, path string, bucket *s3.Bucket) error {
	err := bucket.Del(path)
	fmt.Fprintf(ctx.IO().Progress(), "Deleted: %s\n", path)
	return err
}
