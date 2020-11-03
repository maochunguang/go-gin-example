package tool

import (
	"context"
	"errors"
	"runtime/debug"
	"sync/atomic"
)

type Panic struct {
	R     interface{}
	Stack []byte
}

type PanicGroup struct {
	panics chan Panic // 协程 panic 通知信道
	done   chan int   // 协程完成通知信道
	jobN   int32      // 协程并发数量
}

func NewPanicGroup() *PanicGroup {
	return &PanicGroup{
		panics: make(chan Panic, 8),
		done:   make(chan int, 8),
	}
}

func (g *PanicGroup) Go(f func()) *PanicGroup {
	atomic.AddInt32(&g.jobN, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				g.panics <- Panic{
					R:     r,
					Stack: debug.Stack(),
				}
				return
			}
			g.done <- 1
		}()
		f()
	}()

	return g // 方便链式调用
}
func (g *PanicGroup) Wait(ctx context.Context) error {
	for {
		select {
		case <-g.done:
			if atomic.AddInt32(&g.jobN, -1) == 0 {
				return nil
			}
		case p := <-g.panics:
			//panic(p),此处有两种原则，可以抛出，也可以不抛出
			return errors.New(string(p.Stack))
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
