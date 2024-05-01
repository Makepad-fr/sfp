package core

import (
	"io"
	"net"
	"time"
)

type readOnlyConn struct {
	reader io.Reader
}

func (rc readOnlyConn) Read(b []byte) (int, error) {
	return rc.reader.Read(b)
}

func (rc readOnlyConn) Write(b []byte) (int, error) {
	return 0, io.EOF // No writing should occur
}

func (rc readOnlyConn) Close() error {
	return nil
}

func (rc readOnlyConn) LocalAddr() net.Addr {
	return nil
}

func (rc readOnlyConn) RemoteAddr() net.Addr {
	return nil
}

func (rc readOnlyConn) SetDeadline(t time.Time) error {
	return nil
}

func (rc readOnlyConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (rc readOnlyConn) SetWriteDeadline(t time.Time) error {
	return nil
}
