package workerpool

import "sync"

const (
	DefaultNumOfWorker = 10
)

type Response struct {
	Data  interface{}
	Error error
}

type Handler func(input interface{}) Response

type Pool struct {
	numOfWorker int
	handler     Handler
	inputChan   chan interface{}
	outputChan  chan Response
	wg          sync.WaitGroup
}

type Config struct {
	NumOfWorker int
}

func NewPool(cfg Config) *Pool {
	if cfg.NumOfWorker == 0 {
		cfg.NumOfWorker = DefaultNumOfWorker
	}

	return &Pool{
		numOfWorker: cfg.NumOfWorker,
		inputChan:   make(chan interface{}, cfg.NumOfWorker),
		outputChan:  make(chan Response, cfg.NumOfWorker),
	}
}

func (p *Pool) Run() (chan<- interface{}, chan Response) {
	for i := 0; i < p.numOfWorker; i++ {
		go func() {
			p.wg.Add(1)
			defer p.wg.Done()

			for data := range p.inputChan {
				p.outputChan <- p.handler(data)
			}
		}()
	}

	return p.inputChan, p.outputChan
}

func (p *Pool) Finish() {
	close(p.inputChan)
	p.wg.Wait()
	close(p.outputChan)
}

func (p *Pool) startWorker() {

}

func (p *Pool) RegisterHandler(handler Handler) {
	p.handler = handler
}
