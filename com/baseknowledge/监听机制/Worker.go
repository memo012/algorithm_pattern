package main

import (
	"fmt"
	"time"
)

type Worker struct {
	id  int
	err error
}

func (wk *Worker) worker(workerChan chan<- *Worker) (err error) {
	defer func() {
		//捕获异常信息，防止panic直接退出
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				wk.err = err
			} else {
				wk.err = fmt.Errorf("Panic happened with [%v]", r)
			}
		} else {
			wk.err = err
		}

		//通知 主 Goroutine，当前子Goroutine已经死亡
		workerChan <- wk
	}()

	// do something
	fmt.Println("Start Worker...ID = ", wk.id)

	// 每个worker睡眠一定时间之后，panic退出或者 Goexit()退出
	time.Sleep(time.Second * 5)
	panic("worker panic..")
	//runtime.Goexit()

	return err
}
