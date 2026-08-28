// Copyright 2017-2025 The sacloud/usacloud Authors
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

package common

import (
	"fmt"

	"github.com/sacloud/sacloud-sdk-go/api/simplemq"
	"github.com/sacloud/usacloud/pkg/cli"
)

func NewQueueOp(ctx cli.Context) (simplemq.QueueAPI, error) {
	sdkClient, err := ctx.SDKClient()
	if err != nil {
		return nil, fmt.Errorf("create SDK client: %w", err)
	}
	client, err := simplemq.NewQueueClient(sdkClient)
	if err != nil {
		return nil, fmt.Errorf("create SimpleMQ queue client: %w", err)
	}
	return simplemq.NewQueueOp(client), nil
}

func NewMessageOp(ctx cli.Context, apiKey, queueName string) (simplemq.MessageAPI, error) {
	sdkClient, err := ctx.SDKClient()
	if err != nil {
		return nil, fmt.Errorf("create SDK client: %w", err)
	}
	client, err := simplemq.NewMessageClient(apiKey, sdkClient)
	if err != nil {
		return nil, fmt.Errorf("create SimpleMQ message client: %w", err)
	}
	return simplemq.NewMessageOp(client, queueName), nil
}
