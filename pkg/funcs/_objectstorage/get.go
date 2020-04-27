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

package objectstorage

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func Get(ctx cli.Context, params *params.GetObjectStorageParam) error {
	if ctx.NArgs() > 2 {
		return fmt.Errorf("ObjectStorageGet is failed: %s", "Only two argument can be specified")
	}

	// validate filepath
	filePath := ""
	if ctx.NArgs() > 1 {
		filePath = filepath.Clean(ctx.Args()[1])
	}

	// validate remote path
	if ctx.NArgs() == 0 {
		return fmt.Errorf("<remote path> arg is required")
	}

	// remote path
	path := ""
	path = ctx.Args()[0]
	if path != "" && strings.HasPrefix(path, "/") {
		path = strings.Replace(path, "/", "", 1)
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

	key, _ := bucket.GetKey(path) // if path is directory , GetKey(path) returns nil(with 404 err)
	if key == nil {
		if !params.Recursive {
			return fmt.Errorf("%q is directory. Use -r or --recursive flag", path)
		}

		// path is directory
		if path != "" && !strings.HasSuffix(path, "/") {
			path = path + "/"
		}
		// when path is directory , required local path
		if filePath == "" {
			return fmt.Errorf("<local file/directory> arg is required if it is not a single file download")
		}
		return getRecursive(ctx, path, path, filePath, params.Recursive, bucket)
	}
	// path is file
	return get(ctx, path, filePath, bucket)
}

func getRecursive(ctx cli.Context, remoteBase, remotePath, localBase string, rec bool, bucket *s3.Bucket) error {

	// base: dir1/ , remote: dir1/dir2 -> [localPath]/dir2/
	dirTokens := []string{localBase}
	dirTokens = append(dirTokens, strings.Split(strings.Replace(remotePath, remoteBase, "", 1), "/")...)
	localPath := filepath.Join(dirTokens...)

	// mkdir
	_, err := os.Stat(localPath)
	if err != nil {
		err = os.MkdirAll(localPath, 0755)
		if err != nil {
			return err
		}
	}

	res, err := bucket.List(remotePath, "/", "", 0)
	if err != nil {
		return err
	}

	// first, download files
	for _, content := range res.Contents {
		name := path.Base(content.Key)
		err := get(ctx, content.Key, filepath.Join(localPath, name), bucket)
		if err != nil {
			return err
		}
	}

	if !rec {
		return nil
	}

	// next, download each dir
	for _, pref := range res.CommonPrefixes {
		err := getRecursive(ctx, remoteBase, pref, localBase, rec, bucket)
		if err != nil {
			return nil
		}
	}

	return nil
}

func get(ctx cli.Context, remotePath, localPath string, bucket *s3.Bucket) error {

	// get key
	data, err := bucket.Get(remotePath)
	if err != nil {
		return err
	}

	// write
	var w io.Writer
	if localPath == "" {
		w = ctx.IO().Out()
	} else {
		f, err := os.Create(localPath)
		if err != nil {
			return err
		}
		defer f.Close()
		w = f
	}

	_, err = w.Write(data)
	return err
}
