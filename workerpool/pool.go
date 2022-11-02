package workerpool

import (
	"sync"
	"sync/atomic"
	"time"
)

const (
	DefaultIdleWorker = 10
	DefaultMaxWorker  = 100
	DefaultTimeout    = 2
)

type Handler func()

type List struct {
	queue []Handler
	lock  sync.Mutex
}

type Pool struct {
	maxWorker  int
	idleWorker int

	taskQueue    chan Handler
	workerQueue  chan Handler
	waitingQueue List
	numOfWaiting int32

	stopChan    chan struct{}
	idleTimeout int
	wg          sync.WaitGroup
	lock        sync.Mutex
	stopOnce    sync.Once
	wait        bool
}

type Config struct {
	MaxWorker   int
	IdleWorker  int
	IdleTimeout int
}

func NewPool(cfg Config) *Pool {
	//if cfg.IdleWorker == 0 {
	//	cfg.IdleWorker = DefaultIdleWorker
	//}
	//if cfg.MaxWorker == 0 {
	//	cfg.MaxWorker = DefaultMaxWorker
	//}
	if cfg.IdleTimeout == 0 {
		cfg.IdleTimeout = DefaultTimeout
	}

	return &Pool{
		idleWorker:  cfg.IdleWorker,
		maxWorker:   cfg.MaxWorker,
		idleTimeout: cfg.IdleTimeout,

		taskQueue:   make(chan Handler),
		workerQueue: make(chan Handler),
		waitingQueue: List{
			queue: make([]Handler, 0),
			lock:  sync.Mutex{},
		},
		stopChan: make(chan struct{}),
	}
}

func (p *Pool) Run() {
	defer close(p.stopChan)
	numOfWorker := 0
	idle := false
	timer := time.NewTimer(time.Duration(p.idleTimeout) * time.Second)
Loop:
	for {
		if len(p.waitingQueue.queue) > 0 {
			if !p.processWaitingQueue() {
				break Loop
			}
			continue
		}

		select {
		case h, ok := <-p.taskQueue:
			if !ok {
				break Loop
			}
			select {
			case p.workerQueue <- h:
			default:
				if numOfWorker < p.maxWorker {
					p.wg.Add(1)
					go spawnWorker(h, p.workerQueue, &p.wg)
					numOfWorker += 1
				} else {
					p.pushBackToWaitingQueue(h)
					atomic.StoreInt32(&p.numOfWaiting, int32(len(p.waitingQueue.queue)))
				}
			}
			idle = false
		case <-timer.C:
			if idle && numOfWorker > p.idleWorker {
				select {
				case p.workerQueue <- nil:
					numOfWorker -= 1
				default:
					break
				}
			}

			idle = true
			timer.Reset(time.Duration(p.idleTimeout) * time.Second)
		}
	}

	if p.wait {
		p.runRemainQueue()
	}

	for numOfWorker > 0 {
		p.workerQueue <- nil
		numOfWorker -= 1
	}

	p.wg.Wait()
	timer.Stop()
}

func spawnWorker(h Handler, workerQueue chan Handler, wg *sync.WaitGroup) {
	defer wg.Done()
	for h != nil {
		h()
		h = <-workerQueue
	}
}

func (p *Pool) Submit(h Handler) {
	if h != nil {
		p.taskQueue <- h
	}
}

func (p *Pool) processWaitingQueue() bool {
	select {
	case h, ok := <-p.taskQueue:
		if !ok {
			return false
		}
		p.pushBackToWaitingQueue(h)
	case p.workerQueue <- p.getFrontWaitingQueue():
		p.popFrontWaitingQueue()
	}

	atomic.StoreInt32(&p.numOfWaiting, int32(len(p.waitingQueue.queue)))
	return true
}

func (p *Pool) getFrontWaitingQueue() Handler {
	p.waitingQueue.lock.Lock()
	defer p.waitingQueue.lock.Unlock()
	return p.waitingQueue.queue[0]
}

func (p *Pool) popFrontWaitingQueue() {
	p.waitingQueue.lock.Lock()
	defer p.waitingQueue.lock.Unlock()
	p.waitingQueue.queue = p.waitingQueue.queue[1:]
}

func (p *Pool) pushBackToWaitingQueue(h Handler) {
	p.waitingQueue.lock.Lock()
	defer p.waitingQueue.lock.Unlock()
	p.waitingQueue.queue = append(p.waitingQueue.queue, h)
}

func (p *Pool) Finish(wait bool) {
	p.stopOnce.Do(func() {
		p.lock.Lock()
		p.wait = wait
		p.lock.Unlock()

		close(p.taskQueue)
	})
	<-p.stopChan
}

func (p *Pool) runRemainQueue() {
	for p.numOfWaiting > 0 {
		p.workerQueue <- p.getFrontWaitingQueue()
		p.popFrontWaitingQueue()
		atomic.StoreInt32(&p.numOfWaiting, int32(len(p.waitingQueue.queue)))
	}
}
