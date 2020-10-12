package tbcontext

import (
	"errors"
	"sync"
	"time"
)

type TbContext interface {

	// 当TbContext 被取消或者到了deadline 返回一个被关闭的channel
	TbDone() <-chan struct{}

	// 在 channel Done 关闭后 返回TbContext关闭原因
	TbErr() error

	// 返回TbContext是否被取消 以及自动取消时间
	TbDeadline() (deadline time.Time, ok bool)

	// 获取key对应的value
	TbValue(key interface{}) interface{}
}

// 表明该 tbContext 是可取消的
type tbCanceler interface {
	cancel(removeFromParent bool, err error)
	Done() <-chan struct{}
}

type emptyCtx int

func (*emptyCtx) TbDone() <-chan struct{} {
	return nil
}
func (*emptyCtx) TbErr() error {
	return nil
}
func (*emptyCtx) TbDeadline() (deadline time.Time, ok bool) {
	return
}
func (*emptyCtx) TbValue(key interface{}) interface{} {
	return nil
}

var (
	// background 通常用在Main函数中 作为所有Context的根节点
	background = new(emptyCtx)
	// 通常用在并不知道传递什么Context情形 起到占位符作用
	todo = new(emptyCtx)
)

func Background() TbContext {
	return background
}

func TODO() TbContext {
	return todo
}

type cancelTbCtx struct {
	TbContext
	// 保护之后的字段
	mu       sync.Mutex
	done     chan struct{}
	children map[tbCanceler]struct{}
	err      error
}

func (c *cancelTbCtx) Done() <-chan struct{} {
	panic("implement me")
}

func (c *cancelTbCtx) TbDone() <-chan struct{} {
	c.mu.Lock()
	if c.done == nil {
		c.done = make(chan struct{})
	}
	d := c.done
	c.mu.Unlock()
	return d
}

// closedChan is a reusable closed channel.
var closedChan = make(chan struct{})

// cancel 方法就是关闭channel 递归地取消它的所有子节点 从父节点中删除自己
// 通过关闭channel 将取消信号传递给了它的所有子节点
func (c *cancelTbCtx) cancel(removeFromParent bool, err error) {
	if err == nil {
		panic("context: internal error: missing cancel error")
	}
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return // 已经被其他协程取消
	}
	// 给 err 字段赋值
	c.err = err
	// 关闭 channel 通知其他协程
	if c.done == nil {
		c.done = closedChan
	} else {
		close(c.done)
	}

	// 遍历它的所有子节点
	for child := range c.children {
		// 递归地取消所有子节点
		child.cancel(false, err)
	}

	// 将子节点置空
	c.children = nil
	c.mu.Unlock()

	if removeFromParent {
		// 从父节点中移除自己
		removeChild(c.TbContext, c)
	}
}

// 从父节点中移除自己
func removeChild(parent TbContext, child tbCanceler) {

}

// Canceled is the error returned by Context.Err when the context is canceled.
var Canceled = errors.New("context canceled")

type CancelFunc func()

func WithCancel(parent TbContext) (ctx TbContext, cancel CancelFunc) {
	c := newCancelCtx(parent)
	propagateCancel(parent, &c)
	return &c, func() {
		c.cancel(true, Canceled)
	}
}

func propagateCancel(parent TbContext, c *cancelTbCtx) {

}

func newCancelCtx(parent TbContext) cancelTbCtx {
	return cancelTbCtx{TbContext: parent}
}
