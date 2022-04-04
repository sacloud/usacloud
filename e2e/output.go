// Copyright 2017-2022 The Usacloud Authors
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
	"bufio"
	"context"
	"io"
	"log"
	"strings"
	"sync"
	"testing"
	"time"
)

type Output struct {
	oldOutputs []string
	outputs    []string
	mu         sync.Mutex
}

// CollectOutputs 指定のリーダーをスキャンし、結果を出力バッファにコピーし続ける
func (o *Output) CollectOutputs(prefix string, reader io.ReadCloser) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		o.mu.Lock()
		o.outputs = append(o.outputs, prefix+" "+line)
		o.mu.Unlock()
	}
}

// WaitOutput 出力バッファの中に指定の文字が現れるまで待つ
func (o *Output) WaitOutput(substr string, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	doneCh := make(chan error)

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				doneCh <- ctx.Err()
			case <-ticker.C:
				if o.Contains(substr) {
					o.mu.Lock()
					o.oldOutputs = append(o.oldOutputs, o.outputs...)
					o.outputs = []string{}
					doneCh <- nil
					o.mu.Unlock()
				}
			}
		}
	}()

	return <-doneCh
}

// Contains 出力バッファの中に指定の文字が含まれる場合trueを返す
func (o *Output) Contains(substr string) bool {
	o.mu.Lock()
	defer o.mu.Unlock()
	for _, line := range o.outputs {
		if strings.Contains(line, substr) {
			return true
		}
	}
	return false
}

// Body 現在までの出力バッファの内容を返す
func (o *Output) Body() string {
	o.mu.Lock()
	defer o.mu.Unlock()

	var outputs []string
	outputs = append(outputs, o.oldOutputs...)
	outputs = append(outputs, o.outputs...)
	return strings.Join(outputs, "\n")
}

// Output 出力バッファの内容を標準エラーに出力
func (o *Output) Output() {
	log.Println(o.Body())
}

// Fatal 出力バッファの内容を標準エラーに出力した上でテストをFatalさせる
func (o *Output) Fatal(t *testing.T, args ...interface{}) {
	o.Output()
	t.Fatal(args...)
}

// Fatalf 出力バッファの内容を標準エラーに出力した上でテストをFatalさせる
func (o *Output) Fatalf(t *testing.T, format string, args ...interface{}) {
	o.Output()
	t.Fatalf(format, args...)
}
