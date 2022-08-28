package single

import "sync"

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// single defines a class of works, and any of them can be executed with duplicate suppression
type Single struct {
	locker sync.Mutex
	task   map[string]*call
}

// Do executes a task with key, if two tasks with the same key are invoked, the second task will wait
// until the first one finishes and gets the result from it.
func (s *Single) Do(key string, funk func() (interface{}, error)) (interface{}, error) {
	s.locker.Lock()

	if s.task == nil {
		s.task = make(map[string]*call)
	}

	if task, ok := s.task[key]; ok {
		s.locker.Unlock()
		task.wg.Wait()
		return task.val, task.err
	}

	c := new(call)
	c.wg.Add(1)
	s.task[key] = c
	s.locker.Unlock()

	c.val, c.err = funk()
	c.wg.Done()

	s.locker.Lock()
	delete(s.task, key)
	s.locker.Unlock()

	return c.val, c.err
}
