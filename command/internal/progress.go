package internal

import (
	"fmt"
	"github.com/fatih/color"
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
		timeout:     1 * time.Hour,
		duration:    10 * time.Second,
		wg:          sync.WaitGroup{},
	}
}

func (s *ProgressWriter) Start() {

	s.msgChan = make(chan bool)
	s.wg.Add(1)
	go func() {
		s.print(color.New(color.FgHiWhite), fmt.Sprintf("%s\n", s.msgStart))
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

	clr.Fprint(s.out, msg)
}
