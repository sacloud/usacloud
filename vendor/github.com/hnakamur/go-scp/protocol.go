package scp

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

const (
	msgCopyFile       = 'C'
	msgStartDirectory = 'D'
	msgEndDirectory   = 'E'
	msgTime           = 'T'
)

const (
	replyOK         = '\x00'
	replyError      = '\x01'
	replyFatalError = '\x02'
)

type sourceProtocol struct {
	remIn     io.WriteCloser
	remOut    io.Reader
	remReader *bufio.Reader
}

func newSourceProtocol(remIn io.WriteCloser, remOut io.Reader) (*sourceProtocol, error) {
	s := &sourceProtocol{
		remIn:     remIn,
		remOut:    remOut,
		remReader: bufio.NewReader(remOut),
	}

	return s, s.readReply()
}

func (s *sourceProtocol) WriteFile(fileInfo *FileInfo, body io.ReadCloser) error {
	if !fileInfo.modTime.IsZero() || !fileInfo.accessTime.IsZero() {
		err := s.setTime(fileInfo.modTime, fileInfo.accessTime)
		if err != nil {
			return err
		}
	}
	return s.writeFile(fileInfo.mode, fileInfo.size, fileInfo.name, body)
}

func (s *sourceProtocol) StartDirectory(dirInfo *FileInfo) error {
	if !dirInfo.modTime.IsZero() || !dirInfo.accessTime.IsZero() {
		err := s.setTime(dirInfo.modTime, dirInfo.accessTime)
		if err != nil {
			return err
		}
	}
	return s.startDirectory(dirInfo.mode, dirInfo.name)
}

func (s *sourceProtocol) EndDirectory() error {
	return s.endDirectory()
}

func (s *sourceProtocol) setTime(mtime, atime time.Time) error {
	ms, mus := toSecondsAndMicroseconds(mtime)
	as, aus := toSecondsAndMicroseconds(atime)
	_, err := fmt.Fprintf(s.remIn, "%c%d %d %d %d\n", msgTime, ms, mus, as, aus)
	if err != nil {
		return fmt.Errorf("failed to write scp time header: err=%s", err)
	}
	return s.readReply()
}

func toSecondsAndMicroseconds(t time.Time) (seconds int64, microseconds int) {
	rounded := t.Round(time.Microsecond)
	return rounded.Unix(), rounded.Nanosecond() / int(int64(time.Microsecond)/int64(time.Nanosecond))
}

func (s *sourceProtocol) writeFile(mode os.FileMode, length int64, filename string, body io.ReadCloser) error {
	_, err := fmt.Fprintf(s.remIn, "%c%#4o %d %s\n", msgCopyFile, mode&os.ModePerm, length, filepath.Base(filename))
	if err != nil {
		return fmt.Errorf("failed to write scp file header: err=%s", err)
	}
	_, err = io.Copy(s.remIn, body)
	// NOTE: We close body whether or not copy fails and ignore an error from closing body.
	body.Close()
	if err != nil {
		return fmt.Errorf("failed to write scp file body: err=%s", err)
	}
	err = s.readReply()
	if err != nil {
		return err
	}

	_, err = s.remIn.Write([]byte{replyOK})
	if err != nil {
		return fmt.Errorf("failed to write scp replyOK reply: err=%s", err)
	}
	return s.readReply()
}

func (s *sourceProtocol) startDirectory(mode os.FileMode, dirname string) error {
	// length is not used.
	length := 0
	_, err := fmt.Fprintf(s.remIn, "%c%#4o %d %s\n", msgStartDirectory, mode&os.ModePerm, length, filepath.Base(dirname))
	if err != nil {
		return fmt.Errorf("failed to write scp start directory header: err=%s", err)
	}
	return s.readReply()
}

func (s *sourceProtocol) endDirectory() error {
	_, err := fmt.Fprintf(s.remIn, "%c\n", msgEndDirectory)
	if err != nil {
		return fmt.Errorf("failed to write scp end directory header: err=%s", err)
	}
	return s.readReply()
}

func (s *sourceProtocol) readReply() error {
	b, err := s.remReader.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read scp reply type: err=%s", err)
	}
	if b == replyOK {
		return nil
	}
	if b != replyError && b != replyFatalError {
		return fmt.Errorf("unexpected scp reply type: %v", b)
	}
	line, err := s.remReader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read scp reply message: err=%s", err)
	}
	return &protocolError{
		msg:   line,
		fatal: b == replyFatalError,
	}
}

type sinkProtocol struct {
	remIn     io.WriteCloser
	remOut    io.Reader
	remReader *bufio.Reader
}

func newSinkProtocol(remIn io.WriteCloser, remOut io.Reader) (*sinkProtocol, error) {
	s := &sinkProtocol{
		remIn:     remIn,
		remOut:    remOut,
		remReader: bufio.NewReader(remOut),
	}

	err := s.WriteReplyOK()
	if err != nil {
		return nil, err
	}
	return s, nil
}

type timeMsgHeader struct {
	Mtime time.Time
	Atime time.Time
}

type startDirectoryMsgHeader struct {
	Mode os.FileMode
	Name string
}

type endDirectoryMsgHeader struct{}

type fileMsgHeader struct {
	Mode os.FileMode
	Size int64
	Name string
}

type okMsg struct{}

func fromSecondsAndMicroseconds(seconds int64, microseconds int) time.Time {
	return time.Unix(seconds, int64(microseconds)*(int64(time.Microsecond)/int64(time.Nanosecond)))
}

func (s *sinkProtocol) ReadHeaderOrReply() (interface{}, error) {
	b, err := s.remReader.ReadByte()
	if err == io.EOF {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("failed to read scp message type: err=%s", err)
	}
	switch b {
	case msgCopyFile:
		var h fileMsgHeader
		n, err := fmt.Fscanf(s.remReader, "%04o %d %s\n", &h.Mode, &h.Size, &h.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to read scp file message header: err=%s", err)
		}
		if n != 3 {
			return nil, fmt.Errorf("unexpected count in reading file message header: n=%d", 3)
		}

		err = s.WriteReplyOK()
		if err != nil {
			return nil, fmt.Errorf("failed to write scp replyOK reply: err=%s", err)
		}

		return h, nil
	case msgStartDirectory:
		var h startDirectoryMsgHeader
		var dummySize int64
		n, err := fmt.Fscanf(s.remReader, "%04o %d %s\n", &h.Mode, &dummySize, &h.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to read scp start directory message header: err=%s", err)
		}
		if n != 3 {
			return nil, fmt.Errorf("unexpected count in reading start directory message header: n=%d", 3)
		}

		err = s.WriteReplyOK()
		if err != nil {
			return nil, fmt.Errorf("failed to write scp replyOK reply: err=%s", err)
		}

		return h, nil
	case msgEndDirectory:
		_, err := s.remReader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("failed to read scp end directory message: err=%s", err)
		}

		err = s.WriteReplyOK()
		if err != nil {
			return nil, fmt.Errorf("failed to write scp replyOK reply: err=%s", err)
		}

		return endDirectoryMsgHeader{}, nil
	case msgTime:
		var ms int64
		var mus int
		var as int64
		var aus int
		n, err := fmt.Fscanf(s.remReader, "%d %d %d %d\n", &ms, &mus, &as, &aus)
		if err != nil {
			return nil, fmt.Errorf("failed to read scp time message header: err=%s", err)
		}
		if n != 4 {
			return nil, fmt.Errorf("unexpected count in reading time message header: n=%d", 3)
		}

		err = s.WriteReplyOK()
		if err != nil {
			return nil, fmt.Errorf("failed to write scp replyOK reply: err=%s", err)
		}

		h := timeMsgHeader{
			Mtime: fromSecondsAndMicroseconds(ms, mus),
			Atime: fromSecondsAndMicroseconds(as, aus),
		}

		return h, nil
	case replyOK:
		return okMsg{}, nil
	case replyError, replyFatalError:
		line, err := s.remReader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("failed to read scp reply error message: err=%s", err)
		}

		if b == replyError {
			err = s.WriteReplyOK()
			if err != nil {
				return nil, fmt.Errorf("failed to write scp replyOK reply: err=%s", err)
			}
		}

		return nil, &protocolError{
			msg:   line,
			fatal: b == replyFatalError,
		}
	default:
		return nil, fmt.Errorf("invalid scp message type: %v", b)
	}
}

func (s *sinkProtocol) CopyFileBodyTo(h fileMsgHeader, w io.Writer) error {
	lr := io.LimitReader(s.remReader, h.Size)
	n, err := io.Copy(w, lr)
	if err == io.EOF {
		if n != h.Size {
			return fmt.Errorf("unexpected EOF in CopyFileBodyTo: err=%s", err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to write copy file body: err=%s", err)
	}

	err = s.WriteReplyOK()
	if err != nil {
		return fmt.Errorf("failed to write scp replyOK reply: err=%s", err)
	}

	return nil
}

func (s *sinkProtocol) WriteReplyOK() error {
	_, err := s.remIn.Write([]byte{replyOK})
	return err
}
