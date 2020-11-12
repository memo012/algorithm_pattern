package pool

import (
	"sync"
	"sync/atomic"
	"time"
)

type sig struct{}
type f func() error
type Pool struct {
	capacity       int32         // 池容量
	running        int32         // worker goroutine
	expiryDuration time.Duration // worker 过期时长
	workers        []*Worker     // 存放空闲goroutine
	release        chan sig      // 关闭Pool支持通知所有worker退出运行以防goroutine泄漏
	lock           sync.Mutex
	once           sync.Once
}

func NewPool(size int) (*Pool, error) {
	return NewTimingPool(size, DefaultCleanIntervalTime)
}

func NewTimingPool(size, expiry int) (*Pool, error) {
	if size <= 0 {
		return nil, ErrInvalidPoolSize
	}
	if expiry <= 0 {
		return nil, ErrInvalidPoolExpiry
	}
	p := &Pool{
		capacity:       int32(size),
		release:        make(chan sig, 1),
		expiryDuration: time.Duration(expiry) * time.Second,
	}
	// 启动定期清理过期worker任务 独立goroutine运行
	p.monitorAndClear()
	return p, nil
}

func (p *Pool) monitorAndClear() {
	heartbeat := time.NewTicker(p.expiryDuration)
	for range heartbeat.C {
		currentTime := time.Now()
		p.lock.Lock()
		idleWorkers := p.workers
		// Pool已关闭
		if len(idleWorkers) == 0 && p.Running() == 0 && len(p.release) > 0 {
			p.lock.Unlock()
			return
		}
		n := 0
		for i, w := range idleWorkers {
			// worker 空闲时间小于过期时间
			if currentTime.Sub(w.recycleTime) <= p.expiryDuration {
				break
			}
			n = i
			w.task = nil
			idleWorkers[i] = nil
		}
		n++
		// 表示无可用的worker
		if n >= len(idleWorkers) {
			p.workers = idleWorkers[:0]
		} else {
			p.workers = idleWorkers[:n]
		}
		p.lock.Unlock()
	}
}

func (p *Pool) Submit(task f) error {
	// 判断当前Pool是否关闭
	if len(p.release) > 0 {
		return ErrPoolClosed
	}
	// 获取空闲worker goroutine
	w := p.getWorker()
	w.task <- task
	return nil
}

func (p *Pool) getWorker() *Worker {
	// 判断是否有空闲worker
	// 若存在空闲worker 从worker切片尾部取出
	// 若不存在空闲worker 则判断是否大于cap容量
	// 若小于cap容量 则创建goroutine
	// 若大于cap容量 等待
	var worker *Worker
	waiting := false
	p.lock.Lock()
	idleWorkers := p.workers
	n := len(idleWorkers) - 1
	if n < 0 {
		waiting = p.Running() >= p.Cap()
	} else {
		// 当前队列有空闲worker
		worker = idleWorkers[n]
		idleWorkers[n] = nil
		idleWorkers = idleWorkers[:n]
	}
	p.lock.Unlock()
	if waiting {
		// 容量达到上限
		for {
			p.lock.Lock()
			idleWorkers = p.workers
			l := len(idleWorkers) - 1
			if l < 0 {
				p.lock.Unlock()
				continue
			}
			worker = idleWorkers[l]
			idleWorkers[l] = nil
			idleWorkers = idleWorkers[:l]
			p.lock.Unlock()
			break
		}
	} else if worker == nil {
		// worker 数量没达到cap容量 新建goroutine
		worker = &Worker{
			pool: p,
			task: make(chan f, 1),
		}
		worker.run()
		p.incRunning()
	}
	return worker
}
func (p *Pool) incRunning() {
	atomic.AddInt32(&p.running, 1)
}

func (p *Pool) Running() int {
	return int(atomic.LoadInt32(&p.running))
}

func (p *Pool) Cap() int {
	return int(atomic.LoadInt32(&p.capacity))
}

func (p *Pool) decRunning() {
	atomic.AddInt32(&p.running, -1)
}

func (p *Pool) revertWorker(worker *Worker) {
	// 写入回收时间，亦即该worker的最后一次结束运行的时间
	worker.recycleTime = time.Now()
	p.lock.Lock()
	p.workers = append(p.workers, worker)
	p.lock.Unlock()
}
