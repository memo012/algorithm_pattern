package 并发队列

import "sync"

type SliceQueue struct {
	data  []interface{}
	mutex sync.Mutex
}

func New(count int) *SliceQueue {
	return &SliceQueue{
		data: make([]interface{}, 0, count),
	}
}

// Enqueue 把值放在队尾
func (s *SliceQueue) Enqueue(v interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data = append(s.data, v)
}

// Dequeue 移去队头并返回
func (s *SliceQueue) Dequeue() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.data) == 0 {
		s.mutex.Unlock()
		return nil
	}
	v := s.data[0]
	s.data = s.data[1:]
	return v
}
