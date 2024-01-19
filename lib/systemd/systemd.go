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
	return s
}

func (s *SystemDConn) Connect() (error) {
	s.ctx = context.Background()
	conn, err := dbus.NewWithContext(s.ctx)
	s.conn = conn
	return err
}

func (s *SystemDConn) RestartService(name string) (error) {
	s.Connect()

	// mode docs: https://www.freedesktop.org/wiki/Software/systemd/dbus/#methods
	// read under "StartUnit()" talks about mode string
	channel := make(chan string)
	var err error
	var number int
	go func ()  {
		number, err = s.conn.RestartUnitContext(s.ctx, name, "replace", channel)
	}()
	fmt.Printf("Number from Restart Service: %d", number)
	res := <- channel
	fmt.Printf("; Channel: %s\n", res)

	s.Close()

	return err
}

func (s *SystemDConn) Close() {
	s.conn.Close()
}