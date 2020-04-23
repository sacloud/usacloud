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

package progress

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/sacloud/usacloud/pkg/helper/printer"
)

var mutex = sync.Mutex{}

type ProgressWriter struct {
	out         io.Writer
	printer     *printer.Printer
	msgProgress string
	msgStart    string
	msgComplete string
	msgFail     string
	msgChan     chan bool
	timeout     time.Duration
	duration    time.Duration
	elapsed     time.Duration
	wg          sync.WaitGroup
}

func NewProgress(msgProgress, msgPrefix string, out io.Writer, noColor bool) *ProgressWriter {
	return &ProgressWriter{
		out:         out,
		printer:     &printer.Printer{NoColor: noColor},
		msgProgress: msgProgress,
		msgStart:    fmt.Sprintf("%s is started...", msgPrefix),
		msgComplete: fmt.Sprintf("%s is finished", msgPrefix),
		msgFail:     fmt.Sprintf("%s is failed", msgPrefix),
		timeout:     12 * time.Hour,
		duration:    10 * time.Second,
		wg:          sync.WaitGroup{},
	}
}

type ProgressWriterFunc func(chan bool, chan error)

func ExecWithProgress(msgProgress, msgPrefix string, out io.Writer, noColor bool, f ProgressWriterFunc) error {
	spinner := NewProgress(msgProgress, msgPrefix, out, noColor)
	compChan := make(chan bool)
	errChan := make(chan error)

	// TODO closeを呼ぶ

	go func() {
		spinner.Start()
		f(compChan, errChan)
	}()
progress:
	for {
		select {
		case <-compChan:
			spinner.Stop()
			break progress
		case err := <-errChan:
			return err
		}
	}

	return nil
}

func (s *ProgressWriter) Start() {

	s.msgChan = make(chan bool)
	s.wg.Add(1)
	go func() {
		s.print(color.New(), fmt.Sprintf("%s\n", s.msgStart))
		for {
			select {
			case <-time.After(s.duration):
				s.elapsed += s.duration
				s.print(color.New(color.FgWhite), fmt.Sprintf("\t%s (%.fs elapsed)\n", s.msgProgress, s.elapsed.Seconds()))
			case result := <-s.msgChan:
				if result {
					s.print(color.New(color.FgHiGreen), fmt.Sprintf("%s\n", s.msgComplete))
				} else {
					s.print(color.New(color.FgHiRed), fmt.Sprintf("%s\n", s.msgFail))
				}
				s.wg.Done()
				return
			case <-time.After(s.timeout):
				s.Fail(fmt.Errorf("TimeOut"))
			}
		}
	}()
}

func (s *ProgressWriter) Stop() {
	s.msgChan <- true
	s.wg.Wait()
}

func (s *ProgressWriter) Fail(err error) {
	s.msgFail = fmt.Sprintf("%s: %s", s.msgFail, err)
	s.msgChan <- false
}

func (s *ProgressWriter) print(clr *color.Color, msg string) {
	mutex.Lock()
	defer mutex.Unlock()
	s.printer.Fprint(s.out, clr, msg)
}
