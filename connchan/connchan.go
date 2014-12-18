package connchan

import (
	"net"
)

func Create() (<-chan net.Conn, error) {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		return nil, err
	}

	connchan := make(chan net.Conn)
	go func() {
		defer close(connchan)

		for {
			conn, err := ln.Accept()
			if err != nil {
				return
			}

			connchan <- conn
		}
	}()

	return connchan, nil
}
