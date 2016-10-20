package receiver

import "github.com/open-falcon/gateway/receiver/rpc"

func Start() {
	go rpc.StartRpc()
}
