package receiver

import (
	"log"
	"net"
	"time"

	"github.com/baishancloud/octopux-gateway/g"
	"github.com/baishancloud/octopux-gateway/receiver/rpc"
)

type Server struct {
	Rm        *g.ReceiverStatusManager
	rpcsocket *net.TCPListener
}

func (s *Server) Stop() {
	if s.rpcsocket != nil {
		s.rpcsocket.SetDeadline(time.Now())
		s.Rm.Stop()
	}
}

func New() (*Server, error) {
	s := &Server{}
	rln, err := rpc.NewRpcListener()
	if err != nil {
		log.Println("rpc new Listener error:", err)
		return nil, err

	}
	s.rpcsocket = rln
	s.Rm = &g.ReceiverStatusManager{}
	return s, nil
}

func (s *Server) GoServe() {
	s.Rm.Run()
	go rpc.RpcServe(s.rpcsocket)
}

func Start() {
	go rpc.StartRpc()
}
