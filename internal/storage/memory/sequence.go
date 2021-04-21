package memory

import (
	"sync"
)

type Sequence struct {
	mu    sync.RWMutex
	value int64
}

func (s *Sequence) Next() int64 {
	s.mu.Lock()
	s.value++
	s.mu.Unlock()

	return s.value
}

func (s *Sequence) SetValue(value int64) {
	s.mu.Lock()
	s.value = value
	s.mu.Unlock()
}

func (s *Sequence) Value() int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.value
}
