package MQ

import (
	"errors"
	"sync"
	"time"
)

type Broker interface {
	// 消息推送
	// topic: 订阅的主题
	// message: 传递的消息
	publish(topic string, message interface{}) error

	// 消息的订阅
	// topic: 订阅的主题
	// <-chan interface{}：返回对应的channel通道来接收数据
	subscribe(topic string) (<-chan interface{}, error)

	// 取消订阅
	// topic: 订阅的主题
	// sub: 对应的通道
	unsubscribe(topic string, sub <-chan interface{}) error

	// 关闭消息队列
	close()

	// 内部方法 进行广播消息 对推送的消息进行广播 保证每一个订阅者都可以收到
	broadcast(message interface{}, subscribes []chan interface{})

	// capacity：消息队列的容量
	setConditions(capacity int)
}

type BrokerImpl struct {
	// 关闭消息队列
	exit chan bool

	// 消息队列容量
	capacity int

	// key即是topic，其值则是一个切片，chan类型，这里这么做的原因是我们一个topic可以有多个订阅者，所以一个订阅者对应着一个通道
	topics map[string][]chan interface{}

	// 同步锁
	sync.RWMutex
}

func NewBroker() *BrokerImpl {
	return &BrokerImpl{
		exit:   make(chan bool),
		topics: make(map[string][]chan interface{}),
	}
}

func (b *BrokerImpl) setConditions(capacity int) {
	b.capacity = capacity
}

func (b *BrokerImpl) close() {
	select {
	case <-b.exit:
		return
	default:
		close(b.exit)
		b.Lock()
		b.topics = make(map[string][]chan interface{})
		b.Unlock()
	}
	return
}
func (b *BrokerImpl) broadcast(message interface{}, subscribes []chan interface{}) {
	count := len(subscribes)

	concurrency := 1

	switch {
	case count > 1000:
		concurrency = 3
	case count > 100:
		concurrency = 2
	default:
		concurrency = 1
	}

	// 采用Timer 而不是使用time.After 原因：time.After会产生内存泄漏 在计时器触发之前，垃圾回收器不会回收Timer

	// 延迟时间
	idleDuration := 5 * time.Microsecond
	idleTimeout := time.NewTimer(idleDuration)
	defer idleTimeout.Stop()
	pub := func(start int) {
		for j := start; j < count; j += concurrency {
			idleTimeout.Reset(idleDuration)
			select {
			case subscribes[j] <- message:
			case <-idleTimeout.C:
			case <-b.exit:
				return
			}
		}
	}

	for i := 0; i < concurrency; i++ {
		go pub(i)
	}

}

func (b *BrokerImpl) publish(topic string, message interface{}) error {
	select {
	case <-b.exit:
		return errors.New("broker closed")
	default:
	}

	b.RLock()
	sub, ok := b.topics[topic]
	b.RUnlock()

	if !ok {
		return nil
	}

	// 广播
	b.broadcast(message, sub)
	return nil

}

func (b *BrokerImpl) subscribe(topic string) (<-chan interface{}, error) {
	select {
	case <-b.exit:
		return nil, errors.New("broker closed")
	default:
	}

	ch := make(chan interface{}, b.capacity)
	b.Lock()
	b.topics[topic] = append(b.topics[topic], ch)
	b.Unlock()
	return ch, nil
}

func (b *BrokerImpl) unsubscribe(topic string, sub <-chan interface{}) error {
	select {
	case <-b.exit:
		return errors.New("broker closed")
	default:
	}

	b.RLock()
	subscribes, ok := b.topics[topic]
	b.RUnlock()
	if !ok {
		return nil
	}

	b.Lock()
	var newSub []chan interface{}
	for _, subscribe := range subscribes {
		if subscribe == sub {
			continue
		}
		newSub = append(newSub, subscribe)
	}

	b.topics[topic] = newSub
	b.Unlock()
	return nil
}
