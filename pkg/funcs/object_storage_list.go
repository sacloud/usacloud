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
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func ObjectStorageList(ctx cli.Context, params *params.ListObjectStorageParam) error {

	if ctx.NArgs() > 1 {
		return fmt.Errorf("ObjectStorageList is failed: %s", "Only one argument can be specified")
	}

	// set parameters
	path := ""
	if ctx.NArgs() > 0 {
		path = ctx.Args()[0]
	}
	if path != "" && strings.HasPrefix(path, "/") {
		path = strings.Replace(path, "/", "", 1)
	}
	if path != "" && !strings.HasSuffix(path, "/") {
		path = fmt.Sprintf("%s/", path)
	}

	// on SakuraCloud, bucket name is same as AccessKey
	if params.Bucket == "" {
		params.Bucket = params.AccessKey
	}

	auth, err := aws.GetAuth(params.AccessKey, params.SecretKey)
	if err != nil {
		return fmt.Errorf("ObjectStorageList is failed: %s", err)
	}
	client := s3.New(auth, aws.Region{
		Name:       "us-west-2",
		S3Endpoint: "https://b.sakurastorage.jp",
	})

	bucket := client.Bucket(params.Bucket)

	// list all
	res, err := bucket.List(path, "/", "", 0)
	if err != nil {
		return fmt.Errorf("ObjectStorageList is failed: %s", err)
	}

	list := []interface{}{}
	for i := range res.Contents {
		list = append(list, &res.Contents[i])
	}
	for _, pref := range res.CommonPrefixes {
		list = append(list, map[string]string{"Key": pref})
	}

	return ctx.GetOutput().Print(list...)
}
