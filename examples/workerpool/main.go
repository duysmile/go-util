package main

import (
	"github.com/duysmile/go-util/workerpool"
	"log"
	"time"
)

func main() {
	pool := workerpool.NewPool(workerpool.Config{
		IdleWorker: 2,
		MaxWorker:  4,
	})

	go pool.Run()

	for i := 0; i < 10; i++ {
		i := i
		pool.Submit(func() {
			log.Println("handling: ", i)
			time.Sleep(2 * time.Second)
		})
	}

	//time.Sleep(20 * time.Second)
	pool.Finish(true)
	log.Println("Done")
}
