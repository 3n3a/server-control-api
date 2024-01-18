package systemd

import (
	"context"
	"fmt"

	"github.com/coreos/go-systemd/v22/dbus"
)

type SystemDConn struct {
	ctx context.Context

	conn *dbus.Conn
}

func New() SystemDConn {
	s := SystemDConn{}
	s.Connect()
	return s
}

func (s *SystemDConn) Connect() (error) {
	s.ctx = context.Background()
	conn, err := dbus.NewWithContext(s.ctx)
	s.conn = conn
	return err
}

func (s *SystemDConn) RestartService(name string) (error) {
	// mode docs: https://www.freedesktop.org/wiki/Software/systemd/dbus/#methods
	// read under "StartUnit()" talks about mode string
	channel := make(chan string)
	number, err := s.conn.RestartUnitContext(s.ctx, name, "replace", channel)
	fmt.Printf("Number from Restart Service: %d", number)
	// res := <- channel
	// fmt.Println("; Channel: %s", res)
	return err
}

func (s *SystemDConn) Close() {
	s.conn.Close()
}