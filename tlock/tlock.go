package tlock

import "context"

type TLock struct {
	ch chan struct{}
}

func New() TLock {
	return TLock{ch: make(chan struct{}, 1)}
}

func (t TLock) Lock () {
	t.ch<- struct{}{}
}

func (t TLock) LockWithContext(ctx context.Context) bool {
	select {
	case t.ch <- struct{}{}:
		return true
	case <-ctx.Done():
		return false
	}
}

func (t TLock) UnLock() {
	<-t.ch
}

func (t TLock) TryLock() bool {
	select {
	case t.ch <- struct{}{}:
		return true
	default:
		return false
	}
}
