package utils

type Semaphore struct {
	sem chan struct{}
}

func NewSemaphore(maxConcurrentRequests int) *Semaphore {
	return &Semaphore{sem: make(chan struct{}, maxConcurrentRequests)}
}

func (s *Semaphore) Acquire() {
	s.sem <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.sem
}
