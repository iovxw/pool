package pool

func NewPool(maxThread int, f func()) *Pool {
	return &Pool{
		maxThread: maxThread,
		done:      make(chan bool, maxThread),
		exit:      make(chan bool),
		f:         f,
	}
}

type Pool struct {
	maxThread int
	f         func()
	done      chan bool
	exit      chan bool
}

func (p *Pool) Run() {
	go func() {
		for i := 0; i < p.maxThread; i++ {
			go p.run()
		}
		for {
			select {
			case <-p.done:
				go p.run()
			case <-p.exit:
				return
			}
		}
	}()
}

func (p *Pool) run() {
	p.f()
	p.done <- true
}

func (p *Pool) GetMaxThread() int {
	return p.maxThread
}

func (p *Pool) Exit() {
	p.exit <- true
}
