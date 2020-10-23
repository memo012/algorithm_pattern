package msync

import (
	"fmt"
	"github.com/petermattis/goid"
	"sync"
	"sync/atomic"
	"unsafe"
)

type RecursiveMutex struct {
	sync.Mutex
	owner     int64
	recursion int32
}

const (
	mutexLocked      = 1 << iota // 加锁标识位置
	mutexWoken                   // 唤醒标识位置
	mutexStarving                // 锁饥饿标识位置
	mutexWaiterShift = iota      // 标识waiter的起始bit位置
)

// 尝试获取锁
func (r *RecursiveMutex) TryLock() bool {
	// 1. 尝试获取锁
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&r.Mutex)), 0, mutexLocked) {
		return true
	}
	// 2. 如果处于饥饿状态或者唤醒状态或者加锁状态 这次请求不参与
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&r.Mutex)))
	if old&(mutexLocked|mutexWoken|mutexStarving) != 0 {
		return false
	}
	// 3. 尝试在竞争中获取锁
	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&r.Mutex)), old, new)
}

// 获取等待goroutine数量
func (r *RecursiveMutex) Count() int {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&r.Mutex)))
	v = v >> mutexWaiterShift
	v = v + (v & mutexLocked) //再加上锁持有者的数量，0或者1
	return int(v)
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

func (r *RecursiveMutex) Unlock() {
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
func (r *RecursiveMutex) UnlockTwo(token int64) {
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
