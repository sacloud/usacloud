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

// +build !wasm

package commands

import (
	"io/ioutil"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

func newIO() IO {
	io := &cliIO{
		in:       os.Stdin,
		out:      colorable.NewColorableStdout(),
		progress: colorable.NewColorableStderr(),
		err:      colorable.NewColorableStderr(),
	}
	if !(isatty.IsTerminal(os.Stderr.Fd()) || isatty.IsCygwinTerminal(os.Stderr.Fd())) {
		io.progress = ioutil.Discard
	}
	return io
}
