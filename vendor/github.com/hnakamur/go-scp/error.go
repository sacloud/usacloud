package scp

type protocolError struct {
	msg   string
	fatal bool
}

func (e *protocolError) Error() string { return e.msg }
func (e *protocolError) Fatal() bool   { return e.fatal }
