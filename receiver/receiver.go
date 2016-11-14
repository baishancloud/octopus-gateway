package receiver

import "github.com/baishancloud/octopux-gateway/receiver/rpc"

func Start() {
	go rpc.StartRpc()
}
