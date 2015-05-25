package pool

import (
	"errors"
)

var (
	AlRunErr = errors.New("already running")
)

// 新建线程池
// NewPool(总数量, 最大运行数量, 函数)
// 总数量小于等于 -1 时，为无限，需要手动调用Stop()来停止
func NewPool(max, maxThread int, f func()) Pool {
	return Pool{
		max:       max,
		maxThread: maxThread,
		f:         f,
	}
}

type Pool struct {
	max       int
	maxThread int
	f         func()
	stop      bool
	lock      bool
	n         chan bool
	done      chan bool
}

func (p *Pool) Run() error {
	if p.lock {
		return AlRunErr
	}
	p.lock = true

	// 初始化值，防止上一次Run()产生影响
	p.n = make(chan bool, p.maxThread)
	p.done = make(chan bool, 1)
	p.stop = false
	go func() {
		if p.max > 0 {
			for i := 0; i < p.max; i++ {
				if p.stop {
					break
				}
				p.run()
			}
			p.s()
		} else {
			for {
				if p.stop {
					break
				}
				p.run()
			}
		}
	}()

	return nil
}

func (p *Pool) run() {
	p.n <- true
	go func() {
		p.f()
		<-p.n
	}()
}

func (p *Pool) s() {
	p.stop = true
	for i := 0; i < len(p.n); i++ {
		p.n <- true
	}
	p.lock = false
	p.done <- true
}

// 停止继续补充线程
// 并等待已经开启的线程全部执行完毕
func (p *Pool) Stop() {
	go p.s()
}

// 等待线程全部执行完毕
func (p *Pool) Wait() {
	<-p.done
}
