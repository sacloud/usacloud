// Copyright 2017-2023 The sacloud/usacloud Authors
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

package test

import (
	"io"
	"os"
)

type DummyIOValue struct {
	In       *os.File
	Out      io.Writer
	Progress io.Writer
	Err      io.Writer
}

type DummyIO struct {
	DummyValue *DummyIOValue
}

func (io *DummyIO) In() *os.File {
	return io.DummyValue.In
}

func (io *DummyIO) Out() io.Writer {
	return io.DummyValue.Out
}

func (io *DummyIO) Progress() io.Writer {
	return io.DummyValue.Progress
}

func (io *DummyIO) Err() io.Writer {
	return io.DummyValue.Err
}
