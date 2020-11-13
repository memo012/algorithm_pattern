package 互斥锁

import "time"

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	m := &Mutex{make(chan struct{}, 1)}
	m.ch <- struct{}{}
	return m
}

func (m *Mutex) Lock() {
	<-m.ch
}

func (m *Mutex) UnLock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

// 尝试获取锁
func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}
// 加入一个超时设置
func (m *Mutex) LockTimeout(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-m.ch:
		timer.Stop()
		return true
	case <- timer.C:
	}
	return false
}

// 锁是否被持有
func (m *Mutex) IsLocked() bool {
	return len(m.ch) == 0
}


