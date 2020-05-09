package redis

import (
	"bufio"
	"fmt"
	"net"
)

var DefaultAddr = "127.0.0.1:6379"

type Connection struct {
	Addr     string
	Db       int
	Password string
}

func (con *Connection) RawSend(c net.Conn, cmd []byte) (interface{}, error) {
	_, err := c.Write(cmd)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(c)

	data, err := ReadResponse(reader)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (con *Connection) Connect() (c net.Conn, err error) {
	if con.Addr != "" {
		DefaultAddr = con.Addr
	}
	c, err = net.Dial("tcp", DefaultAddr)
	if err != nil {
		return
	}
	if con.Password != "" {
		cmd := fmt.Sprintf("AUTH %s\r\n", con.Password)
		_, err = con.RawSend(c, []byte(cmd))
		if err != nil {
			return
		}
	}

	if con.Db != 0 {
		cmd := fmt.Sprintf("SELECT %d\r\n", con.Db)
		_, err = con.RawSend(c, []byte(cmd))
		if err != nil {
			return
		}
	}

	return
}
