package 队列的最大值

import "container/list"

type MaxQueue struct {
	queue *list.List
	max   *list.List
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue: list.New(),
		max: list.New(),
	}
}

func (this *MaxQueue) Max_value() int {
	if this.max.Len() == 0 {
		return -1
	}
	return this.max.Front().Value.(int)
}

func (this *MaxQueue) Push_back(value int) {
	this.queue.PushBack(value)
	for this.max.Len() != 0 && this.max.Back().Value.(int) < value {
		this.max.Remove(this.max.Back())
	}
	this.max.PushBack(value)
}

func (this *MaxQueue) Pop_front() int {
	if this.max.Len() != 0 && this.queue.Front().Value.(int) == this.max.Front().Value.(int) {
		this.max.Remove(this.max.Front())
	}
	if this.queue.Len() == 0 {
		return -1
	}
	element := this.queue.Front()
	this.queue.Remove(element)
	return element.Value.(int)
}
