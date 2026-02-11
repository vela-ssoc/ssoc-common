package batchqueue

import (
	"sync"
	"time"
)

type Queue[T any] struct {
	items   []T
	mutex   sync.Mutex
	full    int
	timeout time.Duration
	trigger func([]T)
	timer   *time.Timer
}

// NewQueue 新建一个队列，当队列满了或超时会回调 trigger 函数。
// trigger 不能是长期耗时函数也不能在内部再次调用 Queue.Enqueue。
//
// 要求：
// 1. 该队列不可以启动 worker 协程
// 2. 如果队列未满且不再有新生产数据，也不能让数据得不到消费
func NewQueue[T any](full int, timeout time.Duration, trigger func([]T)) *Queue[T] {
	if full <= 0 {
		full = 1
	}
	if timeout <= 0 {
		timeout = 10 * time.Second
	}
	if trigger == nil {
		trigger = func([]T) {} // noop
	}

	return &Queue[T]{
		items:   make([]T, 0, full),
		full:    full,
		timeout: timeout,
		trigger: trigger,
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.items = append(q.items, item)
	if len(q.items) >= q.full {
		q.callback()
	} else if q.timer == nil {
		q.timer = time.AfterFunc(q.timeout, q.ontimeout)
	}
}

func (q *Queue[T]) ontimeout() {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.callback()
}

func (q *Queue[T]) callback() {
	if len(q.items) == 0 {
		return
	}

	q.trigger(q.items)
	q.items = make([]T, 0, q.full)

	if q.timer != nil {
		q.timer.Stop()
		q.timer = nil
	}
}
