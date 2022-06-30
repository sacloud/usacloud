// Copyright 2017-2022 The sacloud/usacloud Authors
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

package e2e

import (
	"context"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

var (
	installTerraformOnce sync.Once
	execPath             string
)

func newTerraform() (*tfexec.Terraform, error) {
	var initErr error
	installTerraformOnce.Do(func() {
		installer := &releases.ExactVersion{
			Product: product.Terraform,
			Version: version.Must(version.NewVersion("1.0.9")),
		}

		installed, err := installer.Install(context.Background())
		if err != nil {
			initErr = err
			return
		}
		execPath = installed
	})
	if initErr != nil {
		return nil, initErr
	}

	return tfexec.NewTerraform(".", execPath)
}

// InitializeWithTerraform カレントディレクトリで`terraform init & terraform apply --auto-approve`を実行し、
// `terraform destroy --auto-approve`するfuncを返す
//
// 各テストコードで`defer InitializeWithTerraform()()`のように利用する。
func InitializeWithTerraform(t *testing.T) func() {
	// setup
	if err := TerraformInit(); err != nil {
		t.Fatal(err)
	}
	if err := TerraformApply(); err != nil {
		t.Fatal(err)
	}

	// teardown
	return func() {
		if err := TerraformDestroy(); err != nil {
			t.Fatal(err)
		}
	}
}

func TerraformInit() error {
	stateFile := "terraform.tfstate"
	if _, err := os.Stat(stateFile); err == nil {
		if err := os.Remove(stateFile); err != nil {
			return err
		}
	}

	tf, err := newTerraform()
	if err != nil {
		return err
	}
	return tf.Init(context.Background())
}

func TerraformApply() error {
	tf, err := newTerraform()
	if err != nil {
		return err
	}
	return tf.Apply(context.Background())
}

func TerraformRefresh() error {
	tf, err := newTerraform()
	if err != nil {
		return err
	}
	return tf.Refresh(context.Background())
}

func TerraformDestroy() error {
	if os.Getenv("SKIP_CLEANUP") != "" {
		log.Println("Cleanup skipped")
		return nil
	}
	tf, err := newTerraform()
	if err != nil {
		return err
	}
	return tf.Destroy(context.Background())
}
