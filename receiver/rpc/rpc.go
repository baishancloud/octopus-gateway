package rpc

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/baishancloud/octopux-gateway/g"
)

func NewRpcListener() (*net.TCPListener, error) {
	if !g.Config().Rpc.Enabled {
		return nil, errors.New("disable rpc")
	}

	addr := g.Config().Rpc.Listen
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatalf("net.ResolveTCPAddr fail: %s", err)
		return nil, err
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("listen %s fail: %s", addr, err)
		return nil, err
	} else {
		log.Println("rpc listening", addr)
	}
	return listener, nil
}

func RpcServe(ln *net.TCPListener) {
	server := rpc.NewServer()
	server.Register(new(Transfer))

	for {
		conn, err := ln.Accept()
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				ln.Close()
				fmt.Println("Stop rpc accepting connections")
				return
			}
			log.Println("listener.Accept occur error:", err)
			continue
		}
		// go rpc.ServeConn(conn)
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func StartRpc() {
	ln, err := NewRpcListener()
	if err != nil {
		return
	}
	RpcServe(ln)
}
