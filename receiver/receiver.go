package receiver

import (
	"github.com/baishancloud/octopux-gateway/receiver/rpc"
	"github.com/baishancloud/octopux-gateway/receiver/socket"
)

func Start() {
	go rpc.StartRpc()
	go socket.StartSocket()
}
