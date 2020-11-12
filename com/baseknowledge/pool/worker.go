package pool

import "time"

type Worker struct {
	pool *Pool
	task chan f // 任务
	recycleTime time.Time
}

func (w *Worker) run() {
	go func() {
		// 循环监听任务队列 一旦有任务立马取出执行
		for f :=range w.task {
			if f == nil {
				w.pool.decRunning()
				return
			}
			f()
			// worker 回收复用
			w.pool.revertWorker(w)
		}
	}()
}


