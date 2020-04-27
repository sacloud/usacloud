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

package cli

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/sacloud/usacloud/pkg/printer"
)

var mutex = sync.Mutex{}

type Progress struct {
	ctx Context

	out     io.Writer
	printer *printer.Printer
	jobName string

	doneCh chan error

	timeout  time.Duration
	duration time.Duration
	elapsed  time.Duration
}

func NewProgress(ctx Context) *Progress {
	jobName := fmt.Sprintf("%s/%s", ctx.ResourceName(), ctx.CommandName())
	return &Progress{
		ctx:      ctx,
		out:      ctx.IO().Progress(),
		printer:  &printer.Printer{NoColor: ctx.Option().NoColor},
		jobName:  jobName,
		timeout:  12 * time.Hour,
		duration: 10 * time.Second,
	}
}

func (p *Progress) Exec(f func() error) error {
	go p.Start()
	defer p.Stop()

	return f()
}

func (p *Progress) msgPrefix() string {
	if p.ctx.ID().IsEmpty() {
		return p.jobName
	}
	return fmt.Sprintf("%s(ID:%s)", p.jobName, p.ctx.ID().String())
}

func (p *Progress) msgStart() string {
	return fmt.Sprintf("%s: started...\n", p.msgPrefix())
}

func (p *Progress) msgProgress() string {
	return fmt.Sprintf("\t%s: %.fs elapsed\n", p.msgPrefix(), p.elapsed.Seconds())
}

func (p *Progress) msgComplete() string {
	return fmt.Sprintf("%s: done\n", p.msgPrefix())
}

func (p *Progress) msgFailed(err error) string {
	return fmt.Sprintf("%s: failed: %s\n", p.msgPrefix(), err)
}

func (p *Progress) Start() {
	p.doneCh = make(chan error)
	defer close(p.doneCh)

	p.print(color.New(color.FgWhite), p.msgStart())
	for {
		select {
		case <-time.Tick(p.duration):
			p.elapsed += p.duration
			p.print(color.New(color.FgWhite), p.msgProgress())
		case err := <-p.doneCh:
			if err != nil {
				p.print(color.New(color.FgHiRed), p.msgFailed(err))
				return
			}
			p.print(color.New(color.FgHiGreen), p.msgComplete())
			return
		case <-time.After(p.timeout):
			p.doneCh <- fmt.Errorf("timed out")
		case <-p.ctx.Done():
			p.doneCh <- fmt.Errorf("canceled: %s", p.ctx.Err())
		}
	}
}

func (p *Progress) Stop() {
	p.doneCh <- nil
}

func (p *Progress) print(clr *color.Color, msg string) {
	mutex.Lock()
	defer mutex.Unlock()
	p.printer.Fprint(p.out, clr, msg)
	p.printer.Fprint(p.out, color.New(color.Reset), "")
}