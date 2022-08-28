package main

import (
	"github.com/duysmile/go-util/single"
	"log"
	"sync"
	"time"
)

func main() {
	hit := make(map[string]int)
	key := "test"

	runFunc := func() (interface{}, error) {
		// sleep to simulate long-time work
		time.Sleep(1 * time.Second)
		hit[key] += 1
		return hit[key], nil
	}

	var wg sync.WaitGroup
	s := &single.Single{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(s *single.Single) {
			log.Println(s.Do(key, runFunc))
			wg.Done()
		}(s)
	}

	wg.Wait()
}
