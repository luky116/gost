package gxsync

import (
	"fmt"
)

const (
	defaultTaskQNumber = 10
	defaultTaskQLen    = 128
)

/////////////////////////////////////////
// Task Pool Options
/////////////////////////////////////////

type TaskPoolOptions struct {
	tQLen      int // task queue length. buffer size per queue
	tQNumber   int // task queue number. number of queue
	tQPoolSize int // task pool size. number of workers
}

func (o *TaskPoolOptions) validate() {
	if o.tQPoolSize < 1 {
		panic(fmt.Sprintf("illegal pool size %d", o.tQPoolSize))
	}

	if o.tQLen < 1 {
		o.tQLen = defaultTaskQLen
	}

	if o.tQNumber < 1 {
		o.tQNumber = defaultTaskQNumber
	}

	if o.tQNumber > o.tQPoolSize {
		o.tQNumber = o.tQPoolSize
	}
}

type TaskPoolOption func(*TaskPoolOptions)

// @size is the task queue pool size
func WithTaskPoolTaskPoolSize(size int) TaskPoolOption {
	return func(o *TaskPoolOptions) {
		o.tQPoolSize = size
	}
}

// @length is the task queue length
func WithTaskPoolTaskQueueLength(length int) TaskPoolOption {
	return func(o *TaskPoolOptions) {
		o.tQLen = length
	}
}

// @number is the task queue number
func WithTaskPoolTaskQueueNumber(number int) TaskPoolOption {
	return func(o *TaskPoolOptions) {
		o.tQNumber = number
	}
}
