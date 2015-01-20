package pool

import (
	"errors"
)

var (
	AlRunErr = errors.New("already running")
)

func NewPool(maxThread int, f func()) Pool {
	return Pool{
		maxThread: maxThread,
		done:      make(chan bool, maxThread),
		allDone:   make(chan bool),
		f:         f,
	}
}

type Pool struct {
	maxThread int
	threadNum int
	f         func()
	stop      bool
	lock      bool
	done      chan bool
	allDone   chan bool
}

func (p *Pool) Run() error {
	if !p.lock {
		p.lock = true
		p.stop = false
		go func() {
			// 开启预定数量线程
			for i := 0; i < p.maxThread; i++ {
				// 处理当预定数量线程还未开启完毕时
				// 接受到的Stop信号
				if !p.stop {
					p.threadNum++
					go p.run()
				} else {
					break
				}
			}

			// 保持线程数量
			for {
				<-p.done
				if !p.stop {
					go p.run()
				} else {
					p.threadNum--
					if p.threadNum == 0 {
						p.allDone <- true
						p.lock = false
					}
				}
			}
		}()

		return nil
	}

	return AlRunErr
}

func (p *Pool) run() {
	p.f()
	p.done <- true
}

// 获取设置的最大线程数
func (p *Pool) GetMaxThread() int {
	return p.maxThread
}

// 停止继续补充线程
// 并等待已经开启的线程全部执行完毕
func (p *Pool) Stop() {
	p.stop = true
	<-p.allDone
}
