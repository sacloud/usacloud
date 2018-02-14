package internal

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/sacloud/usacloud/helper/printer"
	"io"
	"sync"
	"time"
)

var mutex = sync.Mutex{}

type ProgressWriter struct {
	out         io.Writer
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

func NewProgress(msgProgress, msgPrefix string, out io.Writer) *ProgressWriter {
	return &ProgressWriter{
		out:         out,
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

func ExecWithProgress(msgProgress, msgPrefix string, out io.Writer, f ProgressWriterFunc) error {
	spinner := NewProgress(msgProgress, msgPrefix, out)
	compChan := make(chan bool)
	errChan := make(chan error)

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
	printer.Fprint(s.out, clr, msg)
}
