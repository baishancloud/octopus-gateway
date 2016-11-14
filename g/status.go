package g

import "sync"

type ReceiverStatusManager struct {
	sync.WaitGroup
	lock  sync.RWMutex
	isRun bool
}

func NewReceiverStatusManager() *ReceiverStatusManager {
	rsm := &ReceiverStatusManager{}
	rsm.isRun = false
	return rsm
}

func (r *ReceiverStatusManager) IsRun() bool {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.isRun
}

func (r *ReceiverStatusManager) Run() {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.isRun = true

}

func (r *ReceiverStatusManager) Stop() {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.isRun = false
}
