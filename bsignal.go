package bsignal

import "sync"

// BSignal ...
type BSignal struct {
	wg        sync.WaitGroup
	broadCast chan struct{}
	noCopy    struct{}
}

// NewBSignal ...
func NewBSignal() *BSignal {
	return &BSignal{
		broadCast: make(chan struct{}),
	}
}

// BroadCast notify other goroutine begin to do sth.
func (s *BSignal) BroadCast() {
	close(s.broadCast)
}

// Subscribe Subscribe a bSignal in start.
func (s *BSignal) Subscribe() {
	s.wg.Add(1)
}

// Listened if return value true mean the goroutine really begin to do your thing.
func (s *BSignal) Listened() bool {
	select {
	case <-s.broadCast:
		return true
	default:
		return false
	}
}

// Done check the go routine have done its job.
func (s *BSignal) Done() {
	s.wg.Done()
}

// AllDone main routine wait other routine until all job done.
func (s *BSignal) AllDone() {
	s.wg.Wait()
}
