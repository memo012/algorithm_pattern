package msync

import (
	"fmt"
	"github.com/petermattis/goid"
	"sync"
	"sync/atomic"
)

type RecursiveMutex struct {
	sync.Mutex
	owner     int64
	recursion int32
}

// 方案一：借用开源框架获取goroutine
func (r *RecursiveMutex) Lock() {
	gid := goid.Get()
	// 如果当前持有锁的goroutine就是这次调用的goroutine,说明是重入
	if atomic.LoadInt64(&r.owner) == gid {
		r.recursion++
		return
	}
	r.Mutex.Lock()
	// 获得锁的goroutine第一次调用，记录下它的goroutine id,调用次数加1
	atomic.StoreInt64(&r.owner, gid)
	r.recursion = 1
}

func (r *RecursiveMutex) Unlock()  {
	gid := goid.Get()
	if atomic.LoadInt64(&r.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", r.owner, gid))
	}
	r.recursion--
	if r.recursion != 0 {
		return
	}
	atomic.StoreInt64(&r.owner, -1)
	r.Mutex.Unlock()
}

// 方案二：调用者传入token当做goroutine ID
func (r *RecursiveMutex) lockTwo(token int64) {
	if atomic.LoadInt64(&r.owner) == token {
		r.recursion++
		return
	}
	r.Mutex.Lock()
	atomic.StoreInt64(&r.owner, token)
	r.recursion = 1
}
func (r *RecursiveMutex) UnlockTwo(token int64)  {
	if atomic.LoadInt64(&r.owner) != token {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", token, token))
	}
	r.recursion--
	if r.recursion != 0 {
		return
	}
	atomic.StoreInt64(&r.owner, -1)
	r.Mutex.Unlock()
}
