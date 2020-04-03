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

package commands

import (
	"io"
	"os"
)

// IO .
type IO interface {
	In() *os.File
	Out() io.Writer
	Progress() io.Writer
	Err() io.Writer
}

// cliIO CLIの扱うIO
type cliIO struct {
	in       *os.File
	out      io.Writer
	progress io.Writer
	err      io.Writer
}

func (io *cliIO) In() *os.File {
	return io.in
}

func (io *cliIO) Out() io.Writer {
	return io.out
}

func (io *cliIO) Progress() io.Writer {
	return io.progress
}

func (io *cliIO) Err() io.Writer {
	return io.err
}
